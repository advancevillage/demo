//author: richard
package logs

import (
	"bytes"
	"model"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	log  *model.Log
	logOnce sync.Once
	mutex  sync.RWMutex
)

const (
	Byte = 1
	KB	 = Byte * 1024
	MB   = KB * 1024
	MaxBufferSize = 1 * MB
	MinBufferSize = 256 * KB
)

func InitLog(o *model.Log) (err error) {
	logOnce.Do(func() {
		log  = o
		log.File = strings.Trim(log.File, " ")
		log.CacheSizeString = strings.Trim(log.CacheSizeString, " ")
		//parse CacheSizeString
		var x, y = 0,1
		m := log.CacheSizeString[:len(log.CacheSizeString) - len("XB")]
		n := log.CacheSizeString[len(log.CacheSizeString) - len("XB"):]
		x, err = strconv.Atoi(m)
		if err != nil {
			return
		}
		switch n {
		case "KB","kb","Kb","kB":
			y = KB
		case "MB","mb","Mb","mB":
			y = MB
		case "BB", "bb", "Bb", "bB":
			y = Byte
		}
		log.CacheSize = x * y
		//init log cache
		for i := 0; i < log.CacheCount; i++ {
			cache := make([]byte, 0, log.CacheSize)
			buf := bytes.NewBuffer(cache)
			log.Cache = append(log.Cache, *buf)
		}
		log.R = 0
		log.W = 0
	})
	return
}

func Error(msg string) (err error) {
	var i,j,m,n int
	msg = msg + "\n"
	n = len(msg)
	for n > 0 {
		i = j
		free := log.CacheSize - log.Cache[log.W % log.CacheCount].Len()
		if free >= n {
			j += n
			m, err = log.Cache[log.W % log.CacheCount].Write([]byte(msg[i:j]))
			if err != nil {
				break
			}
			n = n - m
		} else {
			j += free
			m, err = log.Cache[log.W % log.CacheCount].Write([]byte(msg[i:j]))
			if err != nil {
				break
			}
			go func() {
				err = persistent()
			}()
			log.W++
			n = n - m
		}
	}
	return
}

func persistent() (err error) {
	mutex.RLock()
	defer mutex.RUnlock()
	f, err := os.OpenFile(log.File, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0644)
	if err != nil {
		return
	}
	_, err = log.Cache[log.R % log.CacheCount].WriteTo(f)
	if err != nil {
		return
	}
	log.Cache[log.R % log.CacheCount].Reset()
	log.R++
	return
}

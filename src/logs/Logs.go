//author: richard
package logs

import (
	"bufio"
	"log"
	"model"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	instance   *model.Log
	logOnce    sync.Once
)

const (
	Byte = 1
	KB	 = Byte * 1024
	MB   = KB * 1024

	LogLevelEmer  = "[emer]"		//系统级紧急
	LogLevelAlrt  = "[alrt]"		//系统级警告
	LogLevelCrit  = "[crit]"		//系统级危险
	LogLevelEror  = "[eror]"		//用户级错误
	LogLevelWarn  = "[warn]"		//用户级警告
	LogLevelInfo  = "[info]"		//用户级重要
	LogLevelDebg  = "[debg]"		//用户级调试

	LogTimestampLength = len("2019/08/05 23:45:28")
	LogLevelLength     = len("[info]")
	LogShortFileLength = 30
	LogBaseInfoLength  = LogLevelLength + LogTimestampLength + LogShortFileLength
)

func InitLog(o *model.Log) (err error) {
	logOnce.Do(func() {
		instance  = o
		instance.FileName = strings.Trim(instance.FileName, " ")
		instance.File, err = os.OpenFile(instance.FileName, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0644)
		if err != nil {
			return
		}
		instance.CacheSizeString = strings.Trim(instance.CacheSizeString, " ")
		//parse CacheSizeString
		length := len(instance.CacheSizeString)
		var x, y, z = 0,1, length - 1
		for ; z >= 0; z-- {
			if instance.CacheSizeString[z] >= '0' && instance.CacheSizeString[z] <= '9' {
				break
			} else {
				continue
			}
		}
		m := instance.CacheSizeString[:z + 1]
		n := instance.CacheSizeString[z + 1:]
		x, err = strconv.Atoi(m)
		if err != nil {
			return
		}
		switch n {
		case "KB","kb":
			y = KB
		case "MB","mb":
			y = MB
		case "Byte", "byte":
			y = Byte
		}
		instance.CacheSize = x * y
		//init log cache
		for i := 0; i < instance.CacheCount; i++ {
			buf := bufio.NewWriterSize(instance.File, instance.CacheSize)
			instance.Cache = append(instance.Cache, buf)
		}
		instance.Index = 0
	})
	return
}

//@brief: 写入缓存或持久化
func write(level string, message string) {
	length := LogBaseInfoLength + len(message)
	for {
		free := instance.Cache[instance.Index % instance.CacheCount].Available()
		if length > free {
			err := instance.Cache[instance.Index % instance.CacheCount].Flush()
			if err != nil {
				message = err.Error()
				level   = LogLevelEmer
				logger := log.New(instance.File, level, log.LstdFlags | log.Lshortfile)
				logger.Println(message)
				break
			}
			instance.Index = (instance.Index + 1) % instance.CacheCount
		} else {
			logger := log.New(instance.Cache[instance.Index % instance.CacheCount], level, log.LstdFlags | log.Lshortfile)
			logger.Println(message)
			break
		}
	}
}

//@brief: error log
func Error(message string) {
	write(LogLevelEror, message)
}

//@brief: warning log
func Warning(message string) {
	write(LogLevelWarn, message)
}

//@brief: debug log
func Debug(message string) {
	write(LogLevelDebg, message)
}

func Info(message string) {
	write(LogLevelInfo, message)
}

func Alert(message string) {
	write(LogLevelAlrt, message)
}

func Critical(message string) {
	write(LogLevelCrit, message)
}

func Emergency(message string) {
	write(LogLevelEmer, message)
}

func Close() {
	_ = instance.Cache[instance.Index % instance.CacheCount].Flush()
	_ = instance.File.Close()
}
//author: richard
package model

import (
	"bufio"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)
//const
const (
	HttpStatusSuccessCode = http.StatusOK
	HttpStatusBadRequestCode = 400
	HttpStatusInternalServerErrorCode = 500

	//error code
	HttpRequestParamErrorCode = 1000

	HttpRequestParamOffetFormatError = "offset = %s is error"
	HttpRequestParamLimitFormatError = "limit = %s is error"

	//database
	DataBaseQueryErrorCode = 1100
)

//route
type Route struct {
	Method 	string 		  `json:"method"`
	Uri 	string 		  `json:"uri"`
	HandlerFunc httprouter.Handle `json:"handler"`
}

type HttpResponseErrors struct {
	Errors []*HttpResponseErrorsContext	`json:"errors"`
}

type HttpResponseErrorsContext struct {
	Code 	 int 	`json:"code,omitempty"`
	Message  string `json:"message,omitempty"`
	MoreInfo string `json:"moreInfo,omitempty"`
}
//database
type Database struct {
	Host 	 string `xml:"host"`
	Port 	 string `xml:"port"`
	User 	 string `xml:"user"`
	Password string `xml:"password"`
	Schema   string `xml:"schema"`
	Charset  string `xml:"charset"`
	Driver   string `xml:"driver"`
	DB 		*gorm.DB`xml:"-"`
}

type DatabaseMaster struct {
	Member  Database `xml:"database"`
}
type DatabaseSlave struct {
	Members []Database `xml:"database"`
}

type Databases struct {
	Master 	 DatabaseMaster  `xml:"master"`
	Slaves   DatabaseSlave   `xml:"slaves"`
}

type Configure struct {
	LogObject       Log 	  `xml:"log"`
	CachesObject    Caches 	  `xml:"caches"`
	DatabasesObject Databases `xml:"databases"`
}

//log
type Log struct {
	CacheSizeString string 	   `xml:"cache_size"`
	FileName        string 	   `xml:"file"`
	CacheCount 		int 	   `xml:"cache_count"`
	Cache      []*bufio.Writer `xml:"-"`
	CacheSize       int 	   `xml:"-"`
	File       		*os.File   `xml:"-"`
	Index 			int 	   `xml:"-"`
}

//redis master-slave
type Cache struct {
	Host 	 string `xml:"host"`
	Port   	 string `xml:"port"`
	Password string `xml:"password"`
	Client  *redis.Client `xml:"-"`
}
type CacheMaster struct {
	Member  Cache   `xml:"cache"`
}
type CacheSlaves struct {
	Members []Cache `xml:"cache"`
}
type Caches struct {
	Master  CacheMaster	`xml:"master"`
	Slaves  CacheSlaves `xml:"slaves"`
}

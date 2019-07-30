//author: richard
package model

import (
	"bytes"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)
//const
const (
	HttpStatusSuccessCode = http.StatusOK
	HttpStatusBadRequestCode = 400
	HttpStatusInternalServerErrorCode = 500

	//error code
	HttpRequestParamErrorCode = 1000

	//database
	DataBaseQuerryErrorCode = 1100
)

//route
type Route struct {
	Method 	string 		  `json:"method"`
	Uri 	string 		  `json:"uri"`
	HandlerFunc httprouter.Handle `json:"handler"`
}

type HttpResponseErrors struct {
	E []*HttpResponseErrorsContext	`json:"errors"`
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
	DatabasesObject Databases `xml:"databases"`
	LogObject Log `xml:"log"`
}

//log
type Log struct {
	CacheSizeString string 	`xml:"cache_size"`
	File 	        string 	`xml:"file"`
	CacheCount int 			`xml:"cache_count"`
	Cache      []bytes.Buffer `xml:"-"`
	CacheSize  int 			  `xml:"-"`
	R		   int 			  `xml:"-"`
	W 		   int 			  `xml:"-"`
}

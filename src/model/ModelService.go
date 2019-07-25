//author: richard
package model

import (
	"encoding/xml"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
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
	Code 	 int 	`json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"moreInfo"`
}
//database
type Database struct {
	XMLName xml.Name `xml:"database"`
	Host 	 string `xml:"host"`
	Port 	 string `xml:"port"`
	User 	 string `xml:"user"`
	Password string `xml:"password"`
	Schema   string `xml:"schema"`
	Charset  string `xml:"charset"`
	Driver   string `xml:"driver"`
	DB 		*gorm.DB `xml:"-"`
}

type DatabaseMaster struct {
	Configure  Database `xml:"database"`
}
type DatabaseSlave struct {
	Configure  []Database `xml:"database"`
}

type Databases struct {
	XMLName  xml.Name   `xml:"databases"`
	Master 	 DatabaseMaster  `xml:"master"`
	Slaves   DatabaseSlave   `xml:"slaves"`
}

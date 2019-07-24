//author: richard
package model

import (
	"encoding/xml"
	hr "github.com/julienschmidt/httprouter"
)

//route
type Route struct {
	Method 	string 		  `json:"method"`
	Uri 	string 		  `json:"uri"`
	HandlerFunc hr.Handle `json:"handler"`
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
}

type DatabaseMaster struct {
	XMLName    xml.Name `xml:"master"`
	Configure  Database `xml:"database"`
}
type DatabaseSlave struct {
	XMLName    xml.Name   `xml:"slaves"`
	Configure  []Database `xml:"database"`
}

type Databases struct {
	XMLName  xml.Name   `xml:"databases"`
	Master 	 DatabaseMaster  `xml:"master"`
	Slaves   DatabaseSlave   `xml:"slaves"`
}

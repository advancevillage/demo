package model

import "encoding/xml"

type Customer struct {
	XMLName  xml.Name   `json:"-"xml:"customer"`
	ID 		  string 	`json:"id"xml:"id"`
}
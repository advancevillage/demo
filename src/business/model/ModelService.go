package model

import "encoding/xml"

type Customer struct {
	XMLName  xml.Name   `json:"-"     xml:"customer"   gorm:"-"`
	ID 		  string 	`json:"id"    xml:"id"         gorm:"column:id;private_key;type:char(32);not null;"`
	Name 	  string 	`json:"name"  xml:"name"       gorm:"column:name;type:char(32);"`
	Phone     string    `json:"phone" xml:"phone"      gorm:"column:id;type:char(16);"`
}
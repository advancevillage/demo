//author: richard
package router

import hr "github.com/julienschmidt/httprouter"

type Route struct {
	Method 	string 		 `json:"method"`
	Uri 	string 		 `json:"uri"`
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
//author: richard
package router

import hr "github.com/julienschmidt/httprouter"

type Route struct {
	Method 	string 		 `json:"method"`
	Uri 	string 		 `json:"uri"`
	HandlerFunc hr.Handle `json:"handler"`
}

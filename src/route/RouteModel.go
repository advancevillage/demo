//author: richard
package route

import "github.com/gin-gonic/gin"

type Route struct {
	Method 	string 		`json:"method"`
	Uri 	string 		`json:"uri"`
	HandlerFunc gin.HandlerFunc `json:"handler"`
}

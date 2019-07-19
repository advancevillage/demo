//author: richard
package router

import (
	"args"
	hr "github.com/julienschmidt/httprouter"
	"net/http"
)

func RoutePolicy() []*Route {
	var policy = []*Route {
		{Method: "GET", Uri:"/customers", HandlerFunc: QueryCustomersProcessor},
	}
	return policy
}

func NewRouter() (err error) {
	policy := RoutePolicy()
	router := hr.New()
	for i := range policy {
		router.Handle(policy[i].Method, policy[i].Uri, policy[i].HandlerFunc)
	}
	err = http.ListenAndServe(args.HttpServerAddress(), router)
	return
}

// @Summary 查询消费者列表
// @Description 查询消费者列表
// @Produce  json
// @Param offset query int false "页码 >= 1" default(1)
// @Param limit  query int false "每页显示条数 > 0" default(10)
// @Success 200 {string} message " "
// @Header  200 {string} message " "
// @Failure 400 {string} message " "
// @Failure 401 {string} message " "
// @Failure 404 {string} message " "
// @Failure 500 {string} message " "
// @Router /customers [get]
func QueryCustomersProcessor(w http.ResponseWriter, r *http.Request, ps hr.Params) {

}


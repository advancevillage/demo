//author: richard
package router

import (
	"args"
	"business/processor"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"model"
	"net/http"
	"util"
)

func RoutePolicy() []*model.Route {
	var policy = []*model.Route {
		{Method: "GET", Uri:"/v1/customers", HandlerFunc: GetV1QueryCustomersProcessor},
	}
	return policy
}

func NewRouter() (err error) {
	policy := RoutePolicy()
	router := httprouter.New()
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
// @Success 200 {array}  model.Customer
// @Failure 400 {object} model.HttpResponseErrors
// @Failure 401 {object} model.HttpResponseErrors
// @Failure 403 {object} model.HttpResponseErrors
// @Failure 500 {object} model.HttpResponseErrors
// @Router /customers [get]
func GetV1QueryCustomersProcessor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	offset := util.HttpQueryParamsService(r, "offset")
	limit  := util.HttpQueryParamsService(r, "limit")
	//错误信息数据生命期贯穿请求全局
	E := new(model.HttpResponseErrors)
	data, statusCode := processor.QueryCustomersService(offset, limit, E)
	if statusCode == http.StatusOK {
		w.WriteHeader(statusCode)
		n, err := w.Write(data)
		fmt.Println(n, err)
	} else {
		data, err := json.Marshal(E)
		if err != nil {
			fmt.Println(data, err)
		}
		w.WriteHeader(statusCode)
		n, err := w.Write(data)
		fmt.Println(n, err)
	}
}


//author: richard
package router

import (
	"args"
	"business/processor"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"logs"
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
	var err error
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	offset := util.HttpQueryParamsService(r, "offset")
	limit  := util.HttpQueryParamsService(r, "limit")
	//错误信息数据生命期贯穿请求全局
	httpErrorObject := new(model.HttpResponseErrors)
	data, statusCode := processor.QueryCustomersService(offset, limit, httpErrorObject)
	//返回数据集
	w.WriteHeader(statusCode)
	n, err := w.Write(data)
	if err != nil {
		logs.Error(fmt.Sprintf("[%d:%s]", n, err.Error()))
	}
}


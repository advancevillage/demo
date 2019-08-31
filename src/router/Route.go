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
	"net/http/pprof"
	"util"
)

func RoutePolicy() []*model.Route {
	var policy = []*model.Route {
			{Method: "GET", Uri:"/v1/customers", HandlerFunc: GetV1QueryCustomersProcessor},
			{Method: "GET", Uri: "/debug/pprof/", HandlerFunc: GetDebugPProfProcessor},
			{Method: "GET", Uri: "/debug/pprof/allocs", HandlerFunc: GetDebugPProfAllocsProcessor},
			{Method: "GET", Uri: "/debug/pprof/block", HandlerFunc: GetDebugPProfBlockProcessor},
			{Method: "GET", Uri: "/debug/pprof/cmdline", HandlerFunc: GetDebugPProfCmdLineProcessor},
			{Method: "GET", Uri: "/debug/pprof/goroutine", HandlerFunc: GetDebugPProfGoroutineProcessor},
			{Method: "GET", Uri: "/debug/pprof/heap", HandlerFunc: GetDebugPProfHeapProcessor},
			{Method: "GET", Uri: "/debug/pprof/mutex", HandlerFunc: GetDebugPProfMutexProcessor},
			{Method: "GET", Uri: "/debug/pprof/profile", HandlerFunc: GetDebugPProfProfileProcessor},
			{Method: "GET", Uri: "/debug/pprof/symbol",  HandlerFunc: GetDebugPProfSymbolProcessor},
			{Method: "GET", Uri: "/debug/pprof/threadcreate",  HandlerFunc: GetDebugPProfThreadCreateProcessor},
			{Method: "GET", Uri: "/debug/pprof/trace", HandlerFunc: GetDebugPProfTraceProcessor},
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
// @Param offset query int false "页码 >= 0" default(0)
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
	offset := util.HttpQueryParams(r, "offset")
	limit  := util.HttpQueryParams(r, "limit")
	//错误信息数据生命期贯穿请求全局
	httpErrorObject := new(model.HttpResponseErrors)
	data, statusCode := processor.QueryCustomersProcessor(offset, limit, httpErrorObject)
	//返回数据集
	w.WriteHeader(statusCode)
	n, err := w.Write(data)
	if err != nil {
		logs.Error(fmt.Sprintf("[%d:%s]", n, err.Error()))
	}
}

func GetDebugPProfProcessor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pprof.Index(w, r)
}

func GetDebugPProfCmdLineProcessor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pprof.Cmdline(w, r)
}

func GetDebugPProfProfileProcessor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pprof.Profile(w, r)
}

func GetDebugPProfSymbolProcessor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pprof.Symbol(w, r)
}

func GetDebugPProfTraceProcessor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pprof.Trace(w, r)
}

func GetDebugPProfAllocsProcessor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pprof.Handler("allocs").ServeHTTP(w, r)
}

func GetDebugPProfBlockProcessor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pprof.Handler("block").ServeHTTP(w, r)
}

func GetDebugPProfGoroutineProcessor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pprof.Handler("goroutine").ServeHTTP(w, r)
}

func GetDebugPProfHeapProcessor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pprof.Handler("heap").ServeHTTP(w, r)
}

func GetDebugPProfMutexProcessor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pprof.Handler("mutex").ServeHTTP(w, r)
}

func GetDebugPProfThreadCreateProcessor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pprof.Handler("threadcreate").ServeHTTP(w, r)
}
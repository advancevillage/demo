//@brief: processor 负责上层是router,下层是service
//负责参数解析校验等模块,业务的核心逻辑层
package processor

import (
	"business/service"
	"encoding/json"
	"fmt"
	"logs"
	"model"
	"strconv"
)

func QueryCustomersProcessor(offsetString, limitString string, httpErrorObject *model.HttpResponseErrors) (buf []byte, statusCode int) {
	//定义返回对象结构
	var err  error
	var filter = make(map[string]interface{})
	defer func() {
		filter["errors"] = httpErrorObject.Errors
		buf, err = json.Marshal(filter)
		if err != nil {
			logs.Error(err.Error())
		}
	}()
	//校验参数
	offset, e1 := strconv.Atoi(offsetString)
	limit,  e2 := strconv.Atoi(limitString)
	if e1 != nil {
		statusCode = model.HttpStatusBadRequestCode
		httpErrorObject.Errors = append(httpErrorObject.Errors, &model.HttpResponseErrorsContext{
			Code:model.HttpRequestParamErrorCode,
			Message: fmt.Sprintf(model.HttpRequestParamOffetFormatError, offsetString),
		})
		return
	}
	if e2 != nil {
		statusCode = model.HttpStatusBadRequestCode
		httpErrorObject.Errors = append(httpErrorObject.Errors, &model.HttpResponseErrorsContext{
			Code:model.HttpRequestParamErrorCode,
			Message: fmt.Sprintf(model.HttpRequestParamLimitFormatError, limitString),
		})
		return
	}
	var customerService = service.NewCustomerService(false)
	customers, total, err := customerService.Customers(offset, limit)
	if err != nil {
		statusCode = model.HttpStatusInternalServerErrorCode
		logs.Error(err.Error())
		return
	}
	filter["total"] = total
	filter["customers"] = customers
	statusCode = model.HttpStatusSuccessCode
	return
}

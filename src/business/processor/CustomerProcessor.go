//@brief: processor 负责上层是router,下层是service
//负责参数解析校验等模块,业务的核心逻辑层
package processor

import (
	"business/service"
	"encoding/json"
	"fmt"
	"model"
	"strconv"
)

func QueryCustomersProcessor(offsetString, limitString string, httpErrorObject *model.HttpResponseErrors) (buf []byte, statusCode int) {
	offset, e1 := strconv.Atoi(offsetString)
	limit,  e2 := strconv.Atoi(limitString)
	if e1 != nil || e2 != nil {
		buf = nil
		statusCode = model.HttpStatusBadRequestCode
		httpErrorObject.Errors = append(httpErrorObject.Errors, &model.HttpResponseErrorsContext{
			Code:model.HttpRequestParamErrorCode,
			Message: fmt.Sprintf("offset = %s or limit = %s has format error", offsetString, limitString),
		})
		return
	}
	var customerService = service.NewCustomerService(false)
	customers, total, e3 := customerService.Customers(offset, limit)
	//filter response
	filter := make(map[string]interface{})
	filter["total"] = total
	filter["customers"] = customers
	filter["errors"] = httpErrorObject.Errors
	buf, e4 := json.Marshal(filter)
	if e3 != nil || e4 != nil {
		buf = nil
		statusCode = model.HttpStatusInternalServerErrorCode
		httpErrorObject.Errors = append(httpErrorObject.Errors, &model.HttpResponseErrorsContext{
			Code:model.DataBaseQueryErrorCode,
			Message: fmt.Sprintf("database query customer fail"),
		})
		return
	}
	statusCode = model.HttpStatusSuccessCode
	return
}

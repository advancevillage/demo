package processor

import (
	"business/repository"
	"encoding/json"
	"fmt"
	"model"
	"pool"
	"strconv"
)

func QueryCustomersService(offsetString, limitString string, httpErrorObject *model.HttpResponseErrors) (buf []byte, statusCode int) {
	offset, e1 := strconv.Atoi(offsetString)
	limit,  e2 := strconv.Atoi(limitString)
	if e1 != nil || e2 != nil {
		buf = nil
		statusCode = model.HttpStatusBadRequestCode
		httpErrorObject.E = append(httpErrorObject.E, &model.HttpResponseErrorsContext{
			Code:model.HttpRequestParamErrorCode,
			Message: fmt.Sprintf("offset = %s or limit = %s has format error", offsetString, limitString),
		})
		return
	}
	var customerService = &repository.CustomerService{Repo:&repository.CustomerDatabaseRepository{DB:pool.DatabaseConnection(false)}}
	customers, total, e3 := customerService.Customers(offset, limit)

	//filter response
	filter := make(map[string]interface{})
	filter["total"] = total
	filter["customers"] = customers
	buf, e4 := json.Marshal(filter)
	if e3 != nil || e4 != nil {
		buf = nil
		statusCode = model.HttpStatusInternalServerErrorCode
		httpErrorObject.E = append(httpErrorObject.E, &model.HttpResponseErrorsContext{
			Code:model.DataBaseQuerryErrorCode,
			Message: fmt.Sprintf("database query customer fail"),
		})
	}
	statusCode = model.HttpStatusSuccessCode
	return
}
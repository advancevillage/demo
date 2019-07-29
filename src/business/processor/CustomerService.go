package processor

import (
	"fmt"
	"model"
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
	fmt.Println(offset, limit)
	return nil, model.HttpStatusSuccessCode
}
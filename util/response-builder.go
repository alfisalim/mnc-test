package util

import "mnc-test/model"

func ResponseBuilder(responseCode int, message string, data interface{}) *model.CustomResponseHTTP {
	return &model.CustomResponseHTTP{
		StatusCode: responseCode,
		Message: message,
		Data: data,
	}
}

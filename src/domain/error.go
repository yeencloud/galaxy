package domain

import (
	"github.com/yeencloud/ServiceCore/serviceError"
	"net/http"
)

var (
	ErrInvalidMethodName = serviceError.ErrorDescription{HttpCode: http.StatusBadRequest, String: "invalid method name"}
	ErrMethodNotFound    = serviceError.ErrorDescription{HttpCode: http.StatusNotFound, String: "Method not found"}
)
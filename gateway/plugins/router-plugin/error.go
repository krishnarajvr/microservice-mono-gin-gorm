package main

import (
	"net/http"

	"github.com/google/uuid"
)

const (
	BAD_REQUEST           = "BAD_REQUEST"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	INVALID_ARGUMENT      = "INVALID_ARGUMENT"
	OUT_OF_RANGE          = "OUT_OF_RANGE"
	UNAUTHENTICATED       = "UNAUTHENTICATED"
	ACCESS_DENIED         = "ACCESS_DENIED"
	NOT_FOUND             = "NOT_FOUND"
	ABORTED               = "ABORTED"
	ALREADY_EXISTS        = "ALREADY_EXISTS"
	RESOURCE_EXHAUSTED    = "RESOURCE_EXHAUSTED"
	CANCELLED             = "CANCELLED"
	DATA_LOSS             = "DATA_LOSS"
	UNKNOWN               = "UNKNOWN"
	NOT_IMPLEMENTED       = "NOT_IMPLEMENTED"
	UNAVAILABLE           = "UNAVAILABLE"
	DEADLINE_EXCEEDED     = "DEADLINE_EXCEEDED"
)

type Response struct {
	Status    int         `json:"status" example:"200"`
	Data      interface{} `json:"data" example:"{data:{products}}"`
	Error     interface{} `json:"error" example:"{}"`
	RequestId string      `json:"requestId" example:"3b6272b9-1ef1-45e0"`
}

type ErrorData struct {
	Code    string        `json:"code" example:"BAD_REQUEST"`
	Message string        `json:"message" example:"Bad Request"`
	Details []ErrorDetail `json:"details,omitempty"`
}

type ErrorDetail struct {
	Code    string `json:"code" example:"Required"`
	Target  string `json:"target" example:"Name"`
	Message string `json:"message" example:"Name field is required"`
}

type Error404 struct {
	Status uint        `json:"status" example:"404"`
	Error  ErrorData   `json:"error"`
	Data   interface{} `json:"data"`
}

func ErrorResponseWitCode(statusCode int, errorData *ErrorData) Response {
	//Todo : Get from context
	requestId := uuid.New()

	return Response{
		Status:    statusCode,
		Data:      "",
		Error:     errorData,
		RequestId: requestId.String(),
	}
}

func BadRequest(errorMessage string) Response {
	if len(errorMessage) == 0 {
		errorMessage = "Access Denied"
	}

	errorData := &ErrorData{
		Code:    ACCESS_DENIED,
		Message: errorMessage,
	}

	return ErrorResponseWitCode(http.StatusForbidden, errorData)
}

func AccessDenied(errorMessage string) Response {
	if len(errorMessage) == 0 {
		errorMessage = "Access Denied"
	}

	errorData := &ErrorData{
		Code:    ACCESS_DENIED,
		Message: errorMessage,
	}

	return ErrorResponseWitCode(http.StatusForbidden, errorData)
}

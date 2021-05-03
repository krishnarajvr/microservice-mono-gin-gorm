package swagdto

type ErrorNotFound struct {
	Code    string `json:"code" example:"NOT_FOUND"`
	Message string `json:"message" example:"Resource not found"`
}

type ErrorBadRequest struct {
	Code    string        `json:"code" example:"BAD_REQUEST"`
	Message string        `json:"message" example:"Bad Request"`
	Details []ErrorDetail `json:"details"`
}

type ErrorAccessDenied struct {
	Code    string `json:"code" example:"ACCESS_DENIED"`
	Message string `json:"message" example:"Access Denied"`
}

type ErrorInternalError struct {
	Code    string `json:"code" example:"INTERNAL_SERVER_ERROR"`
	Message string `json:"message" example:"Internal server error"`
}

type ErrorDetail struct {
	Code    string `json:"code" example:"Required"`
	Target  string `json:"target" example:"Name"`
	Message string `json:"message" example:"Name field is required"`
}

type Error400 struct {
	Status uint            `json:"status" example:"400"`
	Error  ErrorBadRequest `json:"error"`
	Data   interface{}     `json:"data"`
}

type Error404 struct {
	Status uint          `json:"status" example:"404"`
	Error  ErrorNotFound `json:"error"`
	Data   interface{}   `json:"data"`
}

type Error403 struct {
	Status uint              `json:"status" example:"403"`
	Error  ErrorAccessDenied `json:"error"`
	Data   interface{}       `json:"data"`
}

type Error500 struct {
	Status uint               `json:"status" example:"500"`
	Error  ErrorInternalError `json:"error"`
	Data   interface{}        `json:"data"`
}

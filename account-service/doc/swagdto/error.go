package swagdto

type ErrorData struct {
	Code    string        `json:"code" example:"BAD_REQUEST"`
	Message string        `json:"message" example:"Bad Request"`
	Details []ErrorDetail `json:"details"`
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

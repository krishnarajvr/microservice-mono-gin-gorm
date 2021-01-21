package common

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_NOT_EXIST_ENTITY        = 10011
	ERROR_CHECK_EXIST_ENTITY_FAIL = 10012
	ERROR_ADD_ENTITY_FAIL         = 10013
	ERROR_DELETE_ENTITY_FAIL      = 10014
	ERROR_EDIT_ENTITY_FAIL        = 10015
	ERROR_COUNT_ENTITY_FAIL       = 10016
	ERROR_GET_ENTITYS_FAIL        = 10017
	ERROR_GET_ENTITY_FAIL         = 10018
)

const (
	BAD_REQUEST           = "BAD_REQUEST"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
)

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

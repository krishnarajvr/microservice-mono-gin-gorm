package swagdto

import {

	common "github.com/krishnarajvr/micro-common"
}

type Error404 struct {
	Status uint             `json:"status" example:"404"`
	Error  common.ErrorData `json:"error"`
	Data   interface{}      `json:"data"`
}

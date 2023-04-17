package controller

import (
	"wukong/pkg/response"
	"wukong/server/service"
)

// 接口在controller 中实现

var svc = service.NewService()

var respf = func() *response.Response {
	return response.NewResponse()
}

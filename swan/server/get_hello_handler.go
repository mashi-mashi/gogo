package server

import (
	"gogo/server/gen/models"
	"gogo/server/gen/restapi/factory"

	"github.com/go-openapi/runtime/middleware"
)

func GetHello(p factory.GetHelloParams) middleware.Responder {
	return factory.NewGetHelloOK().WithPayload(&models.APIResponse{Type: "hoge", Message: "hello", Code: 200})
}

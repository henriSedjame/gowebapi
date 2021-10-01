package controllers

import (
	"github.com/hsedjame/gowebapi/framework/web"
	"net/http"
)

type IndexController struct {
}

type HelloEndpoint struct {
}

func (h HelloEndpoint) Path() string {
	return "/hello"
}

func (h HelloEndpoint) Handler() web.RequestHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello Maif vie"))
	}
}

func (h HelloEndpoint) HttpMethod() web.HttpMethod {
	return web.GET
}

func (i IndexController) Endpoints() []web.Endpoint {
	return []web.Endpoint{
		HelloEndpoint{},
	}
}

func (i IndexController) ErrorHandler() web.ErrorHandler {
	return nil
}

func (i IndexController) DefaultModel() interface{} {
	return nil
}

func (i IndexController) ModelKey() interface{} {
	return nil
}

func (i IndexController) Path() string {
	return "/index"
}

func (i IndexController) MiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handler.ServeHTTP(writer, request)
	})
}

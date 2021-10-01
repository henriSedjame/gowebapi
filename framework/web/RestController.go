package web

import (
	"net/http"
)

type ErrorHandler func(error, http.ResponseWriter) error
type RequestHandler func(http.ResponseWriter, *http.Request)
type HttpMethod = string

const (
	GET     HttpMethod = "GET"
	POST    HttpMethod = "POST"
	PUT     HttpMethod = "PUT"
	DELETE  HttpMethod = "DELETE"
	OPTIONS HttpMethod = "OPTIONS"
	PATCH   HttpMethod = "PATCH"
	HEAD    HttpMethod = "HEAD"
	CONNECT HttpMethod = "CONNECT"
	TRACE   HttpMethod = "TRACE"
)

type RestController interface {
	Path() string
	Endpoints() []Endpoint
	MiddleWare(handler http.Handler) http.Handler
	ErrorHandler() ErrorHandler
}

type Endpoint interface {
	Path() string
	Handler() RequestHandler
	HttpMethod() HttpMethod
}

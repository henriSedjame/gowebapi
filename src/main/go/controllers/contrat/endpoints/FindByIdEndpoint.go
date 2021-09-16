package endpoints

import (
	"github.com/hsedjame/gowebapi/framework/web"
)

type FindByIdEndpoint struct {
	RHandler web.RequestHandler
}

func (f FindByIdEndpoint) Path() string {
	return "/{id}"
}

func (f FindByIdEndpoint) Handler() web.RequestHandler {
	return f.RHandler
}

func (f FindByIdEndpoint) HttpMethod() web.HttpMethod {
	return web.GET
}


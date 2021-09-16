package endpoints

import (
	"github.com/hsedjame/gowebapi/framework/web"
)

type DeleteEndpoint struct {
	RHandler web.RequestHandler
}

func (f DeleteEndpoint) Path() string {
	return "/{id}"
}

func (f DeleteEndpoint) Handler() web.RequestHandler {
	return f.RHandler
}

func (f DeleteEndpoint) HttpMethod() web.HttpMethod {
	return web.DELETE
}


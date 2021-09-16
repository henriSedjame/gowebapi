package endpoints

import (
	"github.com/hsedjame/gowebapi/framework/web"
)

type CreateEndpoint struct {
	RHandler web.RequestHandler
}

func (f CreateEndpoint) Path() string {
	return ""
}

func (f CreateEndpoint) Handler() web.RequestHandler {
	return f.RHandler
}

func (f CreateEndpoint) HttpMethod() web.HttpMethod {
	return web.POST
}


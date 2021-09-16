package endpoints

import (
	"github.com/hsedjame/gowebapi/framework/web"
)

type FindAllEndpoint struct {
	RHandler web.RequestHandler
}

func (f FindAllEndpoint) Path() string {
	return ""
}

func (f FindAllEndpoint) Handler() web.RequestHandler {
	return f.RHandler
}

func (f FindAllEndpoint) HttpMethod() web.HttpMethod {
	return web.GET
}


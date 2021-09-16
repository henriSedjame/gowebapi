package endpoints

import (
	"github.com/hsedjame/gowebapi/framework/web"
)

type UpdateEndpoint struct {
	RHandler web.RequestHandler
}

func (f UpdateEndpoint) Path() string {
	return ""
}

func (f UpdateEndpoint) Handler() web.RequestHandler {
	return f.RHandler
}

func (f UpdateEndpoint) HttpMethod() web.HttpMethod {
	return web.PUT
}


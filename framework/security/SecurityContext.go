package security

import "net/http"

type Context struct {
	Authentication Authentication
}

type ContextRepository interface {
	Save(*http.Request, *Context)
	Load(r *http.Request) *Context
}

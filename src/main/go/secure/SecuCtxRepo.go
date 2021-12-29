package secure

import (
	"github.com/hsedjame/gowebapi/framework/security"
	"net/http"
)

type SecuCtxRepo struct {
	Manager security.AuthenticationManager
}

func (s SecuCtxRepo) Save(request *http.Request, context *security.Context) {
	panic("implement me")
}

func (s SecuCtxRepo) Load(r *http.Request) *security.Context {
	jwt := r.Header.Get("Authorization")
	authenticate := s.Manager.Authenticate(security.JwtAuthentication{Jwt: jwt})
	return &security.Context{Authentication: authenticate}
}

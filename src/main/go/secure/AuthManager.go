package secure

import (
	"github.com/hsedjame/gowebapi/framework/security"
)

type AuthManager struct {
}

func (a AuthManager) Authenticate(authentication security.Authentication) security.Authentication {
	return authentication
}

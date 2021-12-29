package security

import "net/http"

type Configuration struct {
	AuthenticationManager     AuthenticationManager
	SecurityContextRepository ContextRepository
	Authorize                 AuthorizeRequestSpec
}

func (config Configuration) Apply(r *http.Request) (*Context, bool) {

	if context := config.SecurityContextRepository.Load(r); context != nil {

		return context, config.Authorize.Authorize(r, config.AuthenticationManager.Authenticate(context.Authentication))

	} else {

		return nil, false

	}

}

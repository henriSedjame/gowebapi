package security

type Configuration struct {
	AuthenticationManager     *AuthenticationManager
	SecurityContextRepository ContextRepository
	Authorize                 AuthorizeRequestSpec
}

func (config Configuration) Apply() {

}

package security

type AuthenticationManager interface {
	Authenticate(Authentication) Authentication
}

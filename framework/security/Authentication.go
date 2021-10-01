package security

type GrantedAuthority struct {
	Authority string
}

type Credentials struct {
	Username string
	Password string
}

type Authentication interface {
	GetAuthorities() []GrantedAuthority
	GetPrincipal() string
	GetCredentials() Credentials
	IsAuthenticated() bool
}

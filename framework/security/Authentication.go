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

type JwtAuthentication struct {
	Jwt string
}

func (j JwtAuthentication) GetAuthorities() []GrantedAuthority {
	return []GrantedAuthority{}
}

func (j JwtAuthentication) GetPrincipal() string {
	return j.Jwt
}

func (j JwtAuthentication) GetCredentials() Credentials {
	return Credentials{
		Username: j.Jwt,
		Password: j.Jwt,
	}
}

func (j JwtAuthentication) IsAuthenticated() bool {
	return true
}

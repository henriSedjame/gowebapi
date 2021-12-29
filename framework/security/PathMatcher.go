package security

import (
	"net/http"
	"regexp"
)

type AuthorizeRequestSpec struct {
	PermitAll    []RequestPathMatcher
	DenyAll      []RequestPathMatcher
	HasAuthority map[string]*RequestPathMatcher
}

type RequestPathMatcher struct {
	LinkRegex string
	Method    string
}

func (matcher RequestPathMatcher) Matches(r *http.Request) bool {
	result := true

	if matcher.Method != "" {
		result = matcher.Method == r.Method
	}

	if matcher.LinkRegex != "" {
		reg := regexp.MustCompile(matcher.LinkRegex)
		result = result && reg.MatchString(r.URL.Path)
	}

	return result
}

func (spec AuthorizeRequestSpec) Authorize(r *http.Request, auth Authentication) bool {

	for _, matcher := range spec.PermitAll {
		if matcher.Matches(r) {
			return true
		}
	}

	for _, matcher := range spec.DenyAll {
		if matcher.Matches(r) {
			return false
		}
	}

	if auth.IsAuthenticated() {
		for _, authority := range auth.GetAuthorities() {
			if matcher := spec.HasAuthority[authority.Authority]; matcher != nil {
				if matcher.Matches(r) {
					return true
				}
			}

		}
	}

	return false
}

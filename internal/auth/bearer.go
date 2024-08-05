package auth

import "github.com/google/go-containerregistry/pkg/authn"

// Create an AuthFactory that produces a bearer authenticator
// from a given bearer token.
func NewBearerAuthFactory(token string) AuthFactory {
	return &bearerAuthFactory{token}
}

type bearerAuthFactory struct {
	token string
}

// Create a bearer authenticator.
func (factory *bearerAuthFactory) CreateAuthenticator() authn.Authenticator {
	return &authn.Bearer{Token: factory.token}
}

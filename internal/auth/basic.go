package auth

import "github.com/google/go-containerregistry/pkg/authn"

// Create an AuthFactory that produces a basic authenticator from the
// given username and password.
func NewBasicAuthFactory(username string, password string) AuthFactory {
	return &basicAuthFactory{username, password}
}

type basicAuthFactory struct {
	username string
	password string
}

// Create a basic authenticator.
func (factory *basicAuthFactory) CreateAuthenticator() authn.Authenticator {
	return &authn.Basic{
		Username: factory.username,
		Password: factory.password,
	}
}

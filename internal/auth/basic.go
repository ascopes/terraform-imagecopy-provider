package auth

import "github.com/google/go-containerregistry/pkg/authn"

func NewBasicAuthProvider(username string, password string) AuthenticatorFactory {
	return &basicAuthProvider{username, password}
}

type basicAuthProvider struct {
	username string
	password string
}

func (provider *basicAuthProvider) CreateAuthenticator() authn.Authenticator {
	return &authn.Basic{
		Username: provider.username,
		Password: provider.password,
	}
}

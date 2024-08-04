package auth

import "github.com/google/go-containerregistry/pkg/authn"

func NewAnonymousProvider() AuthenticatorFactory {
	return &anonymousAuthProvider{}
}

type anonymousAuthProvider struct{}

func (*anonymousAuthProvider) CreateAuthenticator() authn.Authenticator {
	return authn.Anonymous
}

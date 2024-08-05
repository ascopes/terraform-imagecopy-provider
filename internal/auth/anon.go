package auth

import "github.com/google/go-containerregistry/pkg/authn"

// Create an AuthFactory that produces an anonymous authenticator.
func NewAnonymousAuthFactory() AuthFactory {
	return &anonymousAuthFactory{}
}

type anonymousAuthFactory struct{}

// Create an anonymous authenticator.
func (*anonymousAuthFactory) CreateAuthenticator() authn.Authenticator {
	return authn.Anonymous
}

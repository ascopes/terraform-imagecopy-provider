package auth

import "github.com/google/go-containerregistry/pkg/authn"

// A factory type that is capable of building authenticators
// for container registries.
type AuthFactory interface {
	// Create an instance of the authenticator to use.
	CreateAuthenticator() authn.Authenticator
}

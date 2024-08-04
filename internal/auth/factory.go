package auth

import "github.com/google/go-containerregistry/pkg/authn"

// Provider of an authenticator used with OCI repositories.
// This allows creation of authenticator objects lazily on an
// as-needed basis, which will later allow integration with IAM
// based authentication methods such as AWS IAM which only
// provide 15-minute long tokens.

type AuthenticatorFactory interface {
	// Create an instance of the authenticator to use.
	CreateAuthenticator() authn.Authenticator
}

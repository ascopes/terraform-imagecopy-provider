package auth

import "github.com/google/go-containerregistry/pkg/authn"

func NewBearerAuthProvider(token string) AuthenticatorFactory {
	return &bearerAuthProvider{token}
}

type bearerAuthProvider struct {
	token string
}

func (provider *bearerAuthProvider) CreateAuthenticator() authn.Authenticator {
	return &authn.Bearer{Token: provider.token}
}

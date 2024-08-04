package auth

import (
	"testing"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/stretchr/testify/assert"
)

func TestAnonymousAuthProviderReturnsExpectedAuthentication(t *testing.T) {
	// Given
	provider := NewAnonymousProvider()

	// When
	auth := provider.CreateAuthenticator()

	// Then
	assert.Same(t, authn.Anonymous, auth)
}

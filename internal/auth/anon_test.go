package auth

import (
	"testing"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/stretchr/testify/assert"
)

func TestAnonymousAuthFactoryReturnsExpectedAuthentication(t *testing.T) {
	// Given
	factory := NewAnonymousAuthFactory()

	// When
	auth := factory.CreateAuthenticator()

	// Then
	assert.Same(t, authn.Anonymous, auth)
}

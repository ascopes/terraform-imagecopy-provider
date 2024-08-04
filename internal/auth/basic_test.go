package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicAuthProviderReturnsExpectedAuthentication(t *testing.T) {
	// Given
	username := "ashley"
	password := "was_here"
	provider := NewBasicAuthProvider(username, password)

	// When
	auth := provider.CreateAuthenticator()
	token, err := auth.Authorization()

	// Then
	assert.Nil(t, err)
	assert.Equal(t, "ashley", token.Username)
	assert.Equal(t, "was_here", token.Password)
}

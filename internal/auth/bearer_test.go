package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBearerAuthProviderReturnExpectedAuthentication(t *testing.T) {
	// Given
	expectedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	provider := NewBearerAuthProvider(expectedToken)

	// When
	auth := provider.CreateAuthenticator()
	actualToken, err := auth.Authorization()

	// Then
	assert.Nil(t, err)
	assert.Equal(t, expectedToken, actualToken.RegistryToken)
}

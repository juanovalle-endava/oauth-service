package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEnv(t *testing.T) {
	cfg := NewEnv()

	assert.Equal(t, "local", cfg.Environment)
	assert.Equal(t, "cert/jwtRS256.key", cfg.PrivateKeyPath)
	assert.Equal(t, "cert/jwtRS256.key.pub", cfg.PublicKeyPath)
	assert.Equal(t, "24h", cfg.TokenDefaultDuration)
}

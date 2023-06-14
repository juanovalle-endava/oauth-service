package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPrivateKey(t *testing.T) {
	tests := []struct {
		path    string
		wantErr bool
	}{
		{
			path:    "../../cert/jwtRS256.key",
			wantErr: false,
		},
		{
			path:    "../../cert/no_key.key",
			wantErr: true,
		},
	}
	for _, test := range tests {
		pubKey, err := GetPrivateKey(test.path)
		if test.wantErr {
			assert.Error(t, err)
			assert.Empty(t, pubKey)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, pubKey)
		}
	}
}

func TestGetPublicKey(t *testing.T) {
	tests := []struct {
		path    string
		wantErr bool
	}{
		{
			path:    "../../cert/jwtRS256.key.pub",
			wantErr: false,
		},
		{
			path:    "../../cert/no_key.key.pub",
			wantErr: true,
		},
	}
	for _, test := range tests {
		pubKey, err := GetPublicKey(test.path)
		if test.wantErr {
			assert.Error(t, err)
			assert.Empty(t, pubKey)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, pubKey)
		}
	}
}

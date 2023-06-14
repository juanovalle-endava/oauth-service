package utils

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPrivateKey(t *testing.T) {
	tests := []struct {
		path    string
		wantErr bool
	}{
		{
			path:    "../mock/mock_jwtRS256.key",
			wantErr: false,
		},
		{
			path:    "../mock/no_key.key",
			wantErr: true,
		},
	}
	for _, test := range tests {
		privKey, err := GetPrivateKey(test.path)
		if test.wantErr {
			assert.Error(t, err)
			assert.Empty(t, privKey)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, privKey)
		}
	}
}

func TestGetPublicKey(t *testing.T) {
	tests := []struct {
		path    string
		wantErr bool
	}{
		{
			path:    "../mock/mock_jwtRS256.key.pub",
			wantErr: false,
		},
		{
			path:    "../mock/no_key.key.pub",
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

func TestAuthMiddleware(t *testing.T) {
	router := gin.New()
	router.GET("/token", AuthMiddleware(gin.Accounts{
		"user": "root",
	}))

	tests := []struct {
		credentials string
		wantErr     bool
	}{
		{
			credentials: "user:root",
			wantErr:     false,
		},
		{
			credentials: "user:invalid",
			wantErr:     true,
		},
	}
	for _, test := range tests {

		req := httptest.NewRequest("GET", "/token", nil)
		req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(test.credentials)))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if test.wantErr {
			assert.Equal(t, http.StatusUnauthorized, w.Code)
		} else {
			assert.Equal(t, http.StatusOK, w.Code)
		}
	}
}

func TestVerifyMiddleware(t *testing.T) {
	router := gin.New()
	router.GET("/introspect", VerifyMiddleware())

	tests := []struct {
		header  string
		token   string
		wantErr bool
	}{
		{
			header:  "Authorization",
			token:   "Bearer valid_token",
			wantErr: false,
		},
		{
			header:  "",
			token:   "",
			wantErr: true,
		},
		{
			header:  "Authorization",
			token:   "invalid",
			wantErr: true,
		},
	}
	for _, test := range tests {

		req := httptest.NewRequest("GET", "/introspect", nil)
		req.Header.Set(test.header, test.token)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if test.wantErr {
			assert.Equal(t, http.StatusUnauthorized, w.Code)
		} else {
			assert.Equal(t, http.StatusOK, w.Code)
		}
	}
}

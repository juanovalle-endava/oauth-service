package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

var ()

func GetPrivateKey(path string) (*rsa.PrivateKey, error) {
	privKeyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privKeyBytes)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

func GetPublicKey(path string) (*rsa.PublicKey, error) {
	pubKeyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		return nil, err
	}
	return pubKey, nil
}

func AuthMiddleware(accounts gin.Accounts) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, password, ok := ctx.Request.BasicAuth()

		if !ok || accounts[username] != password {
			ctx.Header("WWW-Authenticate", "Basic realm=Restricted")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("username", username)

		ctx.Next()
	}
}

func VerifyMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			ctx.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			ctx.Abort()
			return
		}

		ctx.Set("token", authParts[1])

		ctx.Next()
	}
}

func EncodeToBase64URL(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

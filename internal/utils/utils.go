package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

const (
	NoFileFound         = "key err: there was an error trying to retrieve the file, err: %s"
	KeyParseErr         = "key err: there was an error trying to parse the key, err: %s"
	ParseDurationErr    = "token err: there was an error trying to parse the duration of the token, err: %s"
	SingingTokenErr     = "token err: there was an error trying sign the token, err: %s"
	InvalidTokenErr     = "token err: invalid token, err: %s"
	JwtParseErr         = "token err: there was an error trying to parse the token, err: %s"
	InvalidSignatureErr = "token err: token signature is invalid, err: %s"
)

func GetPrivateKey(path string) (*rsa.PrivateKey, error) {
	privKeyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(NoFileFound, err)
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privKeyBytes)
	if err != nil {
		return nil, fmt.Errorf(KeyParseErr, err)
	}

	return privKey, nil
}

func GetPublicKey(path string) (*rsa.PublicKey, error) {
	pubKeyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(NoFileFound, err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		return nil, fmt.Errorf(NoFileFound, err)
	}
	return pubKey, nil
}

func AuthMiddleware(accounts gin.Accounts) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, password, ok := ctx.Request.BasicAuth()

		if !ok || accounts[username] != password {
			ctx.Header("WWW-Authenticate", "Basic realm=Restricted")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "username or password are incorrect, please try again"})
			ctx.Abort()
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

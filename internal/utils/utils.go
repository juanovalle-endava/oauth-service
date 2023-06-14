package utils

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

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

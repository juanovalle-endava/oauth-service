package models

type KeyResponse struct {
	Keys []Key `json:"keys"`
}

type Key struct {
	Kty string `json:"kty"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
}

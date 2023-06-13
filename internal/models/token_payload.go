package models

import (
	"fmt"
	"time"
)

type TokenPayload struct {
	ClientId  string    `json:"clientId"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func (t TokenPayload) Valid() error {
	if time.Now().After(t.ExpiredAt) {
		return fmt.Errorf("token err: token has expired")
	}
	return nil
}

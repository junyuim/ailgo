package core_auth

import "github.com/golang-jwt/jwt/v4"

type AuthClaims struct {
	jwt.RegisteredClaims

	Name     string `json:"name,omitempty"`
	Scope    string `json:"scope,omitempty"`
	ClientId string `json:"client_id,omitempty"`
}

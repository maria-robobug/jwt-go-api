package models

import jwt "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	Admin    bool   `json:"admin"`
	NextStep string `json:"next_step"`
	jwt.StandardClaims
}

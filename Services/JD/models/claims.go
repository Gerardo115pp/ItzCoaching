package models

import (
	"github.com/golang-jwt/jwt"
)

type AppClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type ExpertTokenInfo struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	IsActive    bool   `json:"is_active"`
	IsAvailable bool   `json:"is_available"`
	ExpertType  string `json:"expert"`
}

type ExpertClaims struct {
	ExpertTokenInfo
	Is_admin bool `json:"is_admin"`
	jwt.StandardClaims
}

type AdminTokenInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

type AdminClaims struct {
	AdminTokenInfo
	Is_admin bool `json:"is_admin"`
	jwt.StandardClaims
}

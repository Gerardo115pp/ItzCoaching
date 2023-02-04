package models

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type RecoverySession struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Used  bool   `json:"used"`
}

type RecoverySessionToken struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.StandardClaims
}

func CreateRecoverySession(email string) *RecoverySession {
	token_id := fmt.Sprintf("recovery-%s", uuid.New().String())
	return &RecoverySession{
		ID:    token_id,
		Email: email,
		Used:  false,
	}
}

func (r *RecoverySession) CreateRecoverySessionToken(jwt_secret string) (string, error) {
	claims := RecoverySessionToken{
		Email: r.Email,
		ID:    r.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwt_secret))
}

func ParseRecoverySessionToken(jwt_secret string, token_string string) (*RecoverySessionToken, error) {
	token, err := jwt.ParseWithClaims(token_string, &RecoverySessionToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwt_secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*RecoverySessionToken)
	if !ok {
		return nil, fmt.Errorf("Could not parse claims")
	}

	return claims, nil
}

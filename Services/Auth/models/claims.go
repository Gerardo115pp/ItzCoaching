package models

import (
	app_config "bonhart_auth_service/Config"
	"time"

	"github.com/golang-jwt/jwt"
)

type ExpertClaims struct {
	ExpertTokenInfo
	Is_admin bool `json:"is_admin"`
	jwt.StandardClaims
}

type AdminClaims struct {
	AdminTokenInfo
	Is_admin bool `json:"is_admin"`
	jwt.StandardClaims
}

func CreateExpertToken(expert *Expert) (string, error) {
	claims := ExpertClaims{
		ExpertTokenInfo: *expert.getTokenInfo(),
		Is_admin:        false,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * (time.Hour * 24)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(app_config.JWT_SECRET))
}

func CreateAdminToken(admin *ItzAdmin) (string, error) {
	claims := AdminClaims{
		AdminTokenInfo: *admin.getTokenInfo(),
		Is_admin:       true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * (time.Hour * 24)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(app_config.JWT_SECRET))
}

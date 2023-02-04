package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bonhart_auth_service/server"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

var jwt_secret string

func SetJWTSecret(secret string) {
	if jwt_secret == "" {
		jwt_secret = secret
	} else {
		echo.EchoFatal(fmt.Errorf("jwt secret is already set"))
	}
}

func JWTKey() string {
	return jwt_secret
}

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(orders_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(HomeResponse{
			Message: "Welcome to the home page",
			Status:  true,
		})

	}
}

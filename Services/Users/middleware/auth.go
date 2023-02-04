package middleware

import (
	"context"
	"fmt"
	app_config "itz_customers_service/Config"
	"itz_customers_service/models"
	"net/http"
	"strings"

	"github.com/Gerardo115pp/patriots_lib/echo"
	"github.com/golang-jwt/jwt"
)

var jwt_secret string

func SetJWTSecret(secret string) {
	if jwt_secret == "" {
		jwt_secret = secret
	} else {
		echo.EchoFatal(fmt.Errorf("jwt secret is already set"))
	}
}

var (
	NO_AUTH_NEEDED = map[string][]string{
		"/experts":         {"GET"}, // don't need to check auth for this route with this method
		"/public_profiles": {"GET"}, // don't need to check auth for this route with this method

	}
)

func shouldCheckToken(route string, request_method string) bool {
	if request_method == "OPTIONS" {
		return false
	}

	for no_auth_route, methods := range NO_AUTH_NEEDED {
		if strings.HasPrefix(route, no_auth_route) {
			for _, method := range methods {
				if method == request_method {
					return false
				}
			}
		}
	}

	return true
}

func CheckAuth(next func(response http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Authorization")
		if token == app_config.DOMAIN_SECRET {
			new_request := request.WithContext(context.WithValue(request.Context(), "user_id", app_config.DOMAIN_SECRET))
			next(response, new_request)
			return
		}

		if !shouldCheckToken(request.URL.Path, request.Method) {
			next(response, request)
			return
		}
		echo.Echo(echo.GreenFG, fmt.Sprintf("Checking auth for %s", request.URL.Path))

		if token == "" {
			token = request.URL.Query().Get("token")
		}

		token = strings.TrimPrefix(token, "Bearer ")
		token = strings.TrimSpace(token)
		if token == "" {
			response.WriteHeader(401) // Unauthorized
			return
		}

		token_data, err := jwt.ParseWithClaims(token, &models.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwt_secret), nil
		})

		if err != nil {
			response.WriteHeader(401) // Unauthorized
			return
		}

		var user_id string = token_data.Claims.(*models.AppClaims).UserID
		new_request := request.WithContext(context.WithValue(request.Context(), "user_id", user_id))

		next(response, new_request)
	}
}

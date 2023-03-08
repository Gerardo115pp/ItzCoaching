package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	app_config "itz_JD_service/Config"
	"itz_JD_service/models"
	"net/http"
	"reflect"
	"strings"

	"github.com/Gerardo115pp/patriots_lib/echo"
	"github.com/golang-jwt/jwt"
)

var (
	NO_AUTH_NEEDED = map[string][]string{
		"/appointments/expert": {"GET"},
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

func decodeJWTPayload(token string) (map[string]any, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("Invalid token format")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	var payload_map map[string]any
	err = json.Unmarshal(payload, &payload_map)
	if err != nil {
		return nil, err
	}

	if _, exists := payload_map["id"]; exists {
		if reflect.TypeOf(payload_map["id"]) != reflect.TypeOf(int(0)) {
			switch payload_map["id"].(type) {
			case float64:
				payload_map["id"] = int(payload_map["id"].(float64))
			default:
				return nil, fmt.Errorf("Invalid id type")
			}
		}
	}

	return payload_map, nil
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

		var err error

		payload, err := decodeJWTPayload(token)
		echo.Echo(echo.GreenFG, fmt.Sprintf("payload: %v", payload))
		if err != nil {
			echo.Echo(echo.RedFG, fmt.Sprintf("Error decoding token: %s", err.Error()))
			response.WriteHeader(401) // Unauthorized
			return
		}

		if _, exists := payload["is_admin"]; !exists {
			echo.Echo(echo.RedFG, fmt.Sprintf("Error decoding token: is_admin not found"))
			response.WriteHeader(401) // Unauthorized
			return
		}

		if payload["is_admin"] == true {
			_, err = ValidateAdminAuth(token, response, request)
		} else if payload["is_admin"] == false {
			_, err = ValidateExpertAuth(token, response, request)
		} else {
			err = fmt.Errorf("Invalid token")
		}

		if err != nil {
			echo.Echo(echo.OrangeFG, fmt.Sprintf("bad token: %s", token))
			echo.Echo(echo.RedFG, fmt.Sprintf("Error validating token: %s", err.Error()))
			response.WriteHeader(401) // Unauthorized
			return
		}

		new_request := request.WithContext(context.WithValue(request.Context(), "claim_data", payload))

		next(response, new_request)
	}
}

func ValidateExpertAuth(token string, response http.ResponseWriter, request *http.Request) (user_id int, err error) {
	echo.Echo(echo.GreenFG, fmt.Sprintf("Validating expert auth for %s", request.URL.Path))
	token_data, err := jwt.ParseWithClaims(token, &models.ExpertClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(app_config.JWT_SECRET), nil
	})
	if err != nil {
		response.WriteHeader(401) // Unauthorized
		return
	}

	user_id = token_data.Claims.(*models.ExpertClaims).ID

	return
}

func ValidateAdminAuth(token string, response http.ResponseWriter, request *http.Request) (user_id int, err error) {
	echo.Echo(echo.GreenFG, fmt.Sprintf("Validating admin auth for %s", request.URL.Path))
	token_data, err := jwt.ParseWithClaims(token, &models.AdminClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(app_config.JWT_SECRET), nil
	})
	if err != nil {
		response.WriteHeader(401) // Unauthorized
		return
	}

	user_id = token_data.Claims.(*models.AdminClaims).ID

	return
}

package handlers

import (
	"bonhart_auth_service/models"
	"bonhart_auth_service/repository"
	"bonhart_auth_service/server"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

type AdminAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AdminsHandler(auth_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getAdminsHandler(response, request)
		case http.MethodPost:
			postAdminsHandler(response, request)
		case http.MethodPatch:
			patchAdminsHandler(response, request)
		case http.MethodDelete:
			deleteAdminsHandler(response, request)
		case http.MethodPut:
			putAdminsHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(200) // V8 crys really hard and patheticly when you don't return a status code
		default:
			response.WriteHeader(405)
		}
	}
}

func getAdminsHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Get Requests: returns a expert token which is an authentication jwt for the Expert user  */
	response.WriteHeader(405)
}

// Create a new expert
func postAdminsHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Post Requests */
	var admin_auth *AdminAuth = new(AdminAuth)

	err := json.NewDecoder(request.Body).Decode(admin_auth)
	if err != nil {
		echo.EchoWarn("Error decoding admin auth")
		response.WriteHeader(400)
		return
	}

	target_admin, err := repository.GetAdminByEmail(request.Context(), admin_auth.Email)
	if err != nil {
		echo.EchoWarn("Admin not found")
		response.WriteHeader(404)
		return
	}

	err = target_admin.ComparePassword(admin_auth.Password)

	if err != nil {
		echo.EchoWarn("Wrong password")
		response.WriteHeader(403)
		return
	}

	admin_jwt, err := models.CreateAdminToken(target_admin)
	if err != nil {
		echo.EchoWarn("Error creating admin token")
		response.WriteHeader(500)
		return
	}

	response.WriteHeader(200)
	response.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, admin_jwt)))
}

func patchAdminsHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Patch Requests */
	response.WriteHeader(405) // Method not allowed
}

func deleteAdminsHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Delete Requests */
	response.WriteHeader(405) // Method not allowed
}

func putAdminsHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Put Requests */
	response.WriteHeader(405) // Method not allowed
}

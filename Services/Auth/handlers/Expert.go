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

type ExpertAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ExpertsHandler(auth_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getExpertsHandler(response, request)
		case http.MethodPost:
			postExpertsHandler(response, request)
		case http.MethodPatch:
			patchExpertsHandler(response, request)
		case http.MethodDelete:
			deleteExpertsHandler(response, request)
		case http.MethodPut:
			putExpertsHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(200) // V8 crys really hard and patheticly when you don't return a status code
		default:
			response.WriteHeader(405)
		}
	}
}

func getExpertsHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Get Requests: returns a expert token which is an authentication jwt for the Expert user  */
	response.WriteHeader(405)
}

// Create a new expert
func postExpertsHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Post Requests */
	var expert_auth *ExpertAuth = new(ExpertAuth)

	err := json.NewDecoder(request.Body).Decode(expert_auth)
	if err != nil {
		echo.EchoWarn("Error decoding expert auth")
		response.WriteHeader(400)
		return
	}

	target_expert, err := repository.GetExpertByEmail(request.Context(), expert_auth.Email)
	if err != nil {
		echo.EchoWarn("Expert not found")
		response.WriteHeader(404)
		return
	}

	err = target_expert.ComparePassword(expert_auth.Password)

	if err != nil {
		echo.EchoWarn("Wrong password")
		response.WriteHeader(403)
		return
	}

	expert_jwt, err := models.CreateExpertToken(target_expert)
	if err != nil {
		echo.EchoWarn("Error creating expert token")
		response.WriteHeader(500)
		return
	}

	response.WriteHeader(200)
	response.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, expert_jwt)))
}

func patchExpertsHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Patch Requests */
	response.WriteHeader(405) // Method not allowed
}

func deleteExpertsHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Delete Requests */
	response.WriteHeader(405) // Method not allowed
}

func putExpertsHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Put Requests */
	response.WriteHeader(405) // Method not allowed
}

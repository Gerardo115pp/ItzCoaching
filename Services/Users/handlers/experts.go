package handlers

import (
	"encoding/json"
	"fmt"
	"itz_customers_service/helpers"
	"itz_customers_service/models"
	"itz_customers_service/repository"
	"itz_customers_service/server"
	"net/http"
	"strconv"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

type PostExpertRequest struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ExpertType string `json:"expert_type"`
}

type UpdateExpertRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsActive    bool   `json:"is_active,omitempty"`
	IsAvailable bool   `json:"is_available,omitempty"`
	ExpertType  string `json:"expert_type"`
}

func ExpertsHandler(order_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getExpertHandler(response, request)
		case http.MethodPost:
			postExpertHandler(response, request)
		case http.MethodPatch:
			patchExpertHandler(response, request)
		case http.MethodDelete:
			deleteExpertHandler(response, request)
		case http.MethodPut:
			putExpertHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func putExpertHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func deleteExpertHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func patchExpertHandler(response http.ResponseWriter, request *http.Request) {
	var updated_expert_request *UpdateExpertRequest = new(UpdateExpertRequest)

	err := json.NewDecoder(request.Body).Decode(updated_expert_request)
	if err != nil {
		echo.Echo(echo.OrangeFG, "Expert patch handler error: failed to decode request")
		echo.EchoErr(err)
		response.WriteHeader(400)
		return
	}

	if updated_expert_request.ID == 0 {
		echo.Echo(echo.OrangeFG, "Expert patch handler error: id is required")
		response.WriteHeader(422)
		return
	}

	var target_expert *models.Expert

	target_expert, err = repository.GetExpertByID(request.Context(), updated_expert_request.ID)
	if err != nil {
		echo.Echo(echo.OrangeFG, "Expert patch handler error: failed to get expert")
		echo.EchoErr(err)
		response.WriteHeader(404)
	}

	var raw_request map[string]interface{} = helpers.StructToMap(updated_expert_request)
	target_expert.Update(raw_request)

	fmt.Printf("target_expert: %v\n", target_expert)
	err = repository.UpdateExpert(request.Context(), target_expert)
	if err != nil {
		echo.Echo(echo.OrangeFG, "Expert patch handler error: failed to update expert")
		echo.EchoErr(err)
		response.WriteHeader(500)
		return
	}

	echo.Echo(echo.GreenFG, "Expert patch handler: expert updated successfully")
	response.WriteHeader(200)
	return
}

func postExpertHandler(response http.ResponseWriter, request *http.Request) {
	// var creator_id string = request.Context().Value("user_id").(string) TODO: renable when auth service is ready( and also convert to int)
	var creator_id int = 1
	var new_expert_request *PostExpertRequest = new(PostExpertRequest)

	err := json.NewDecoder(request.Body).Decode(new_expert_request)
	if err != nil {
		echo.EchoErr(err)
		response.WriteHeader(400)
		return
	}

	var raw_request map[string]interface{} = helpers.StructToMap(new_expert_request)
	raw_request["created_by"] = creator_id

	new_expert, err := models.CreateExpertObject(raw_request)

	if err != nil {
		echo.EchoErr(err)
		response.WriteHeader(400)
	}

	err = repository.InsertExpert(request.Context(), new_expert)
	if err != nil {
		echo.Echo(echo.OrangeFG, "Expert post handler error: failed to inserting expert")
		echo.EchoErr(err)
		response.WriteHeader(409)
		return
	}

	response.WriteHeader(201)

	return
}

func getExpertHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("id") != "" {
		getExpertByIDHandler(response, request)
		return
	} else if request.URL.Query().Get("email") != "" {
		getExpertByEmailHandler(response, request)
		return
	} else if request.URL.Query().Get("all") == "1" {
		getAllExpertsHandler(response, request)
		return
	}

	echo.Echo(echo.OrangeFG, fmt.Sprintf("Cannot process request: %s", request.URL.Path))
	response.WriteHeader(501)
	return
}

func getAllExpertsHandler(response http.ResponseWriter, request *http.Request) {
	var experts []*models.Expert
	var err error
	echo.Echo(echo.GreenFG, "Expert get all handler: getting all experts")
	if auth_header := request.Header.Get("Authorization"); auth_header == "" {
		echo.Echo(echo.OrangeFG, "Expert get all handler error: no auth header provided")
		response.WriteHeader(401)
		return
	}

	experts, err = repository.GetAllExperts(request.Context())
	if err != nil {
		echo.Echo(echo.OrangeFG, "Expert get all handler error: failed to get experts")
		echo.EchoErr(err)
		response.WriteHeader(500)
		return
	}

	//clean the password data

	var experts_safedata []map[string]any

	for _, expert := range experts {
		expert_safedata := helpers.StructToMap(expert)
		delete(expert_safedata, "password")
		experts_safedata = append(experts_safedata, expert_safedata)
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(experts_safedata)
}

// CALLED BY: getExpertHandler
func getExpertByEmailHandler(response http.ResponseWriter, request *http.Request) {
	var email string = request.URL.Query().Get("email")
	if email == "" {
		echo.Echo(echo.OrangeFG, "Expert get by email handler error: no email provided")
		response.WriteHeader(400)
		return
	}

	expert, err := repository.GetExpertByEmail(request.Context(), email)
	if err != nil {
		echo.EchoErr(err)
		response.WriteHeader(404)
		return
	}

	var public_data map[string]any = expert.PublicData()

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(public_data)

	return
}

// CALLED BY: getExpertHandler
func getExpertByIDHandler(response http.ResponseWriter, request *http.Request) {
	var expert_id int

	expert_id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		echo.Echo(echo.OrangeFG, fmt.Sprintf("Expert get by id handler error: failed to convert id '%s' to int", request.URL.Query().Get("id")))
		echo.EchoErr(err)
		response.WriteHeader(400)
		return
	}

	expert, err := repository.GetExpertByID(request.Context(), expert_id)
	if err != nil {
		echo.EchoErr(err)
		response.WriteHeader(404)
		return
	}

	var public_data map[string]any = expert.PublicData()

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(public_data)

	return
}

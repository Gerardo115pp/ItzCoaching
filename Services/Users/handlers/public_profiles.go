package handlers

import (
	"encoding/json"
	"fmt"
	"itz_customers_service/models"
	"itz_customers_service/repository"
	"itz_customers_service/server"
	"net/http"
	"strconv"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

func PublicProfilesHandler(order_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getPublicProfileHandler(response, request)
		case http.MethodPost:
			postPublicProfileHandler(response, request)
		case http.MethodPatch:
			patchPublicProfileHandler(response, request)
		case http.MethodDelete:
			deletePublicProfileHandler(response, request)
		case http.MethodPut:
			putPublicProfileHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func putPublicProfileHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func deletePublicProfileHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func patchPublicProfileHandler(response http.ResponseWriter, request *http.Request) {
	var authenicaded_claims map[string]any = request.Context().Value("claim_data").(map[string]any)
	var updated_public_profile *models.PublicProfile = new(models.PublicProfile)

	err := json.NewDecoder(request.Body).Decode(updated_public_profile)
	if err != nil {
		echo.Echo(echo.OrangeFG, "PublicProfile patch handler error: failed to decode request")
		echo.EchoErr(err)
		response.WriteHeader(400)
		return
	}

	if updated_public_profile.Expert_id == 0 {
		echo.Echo(echo.OrangeFG, "PublicProfile patch handler error: id is required")
		response.WriteHeader(422)
		return
	} else if authenicaded_claims["id"].(int) != updated_public_profile.Expert_id && authenicaded_claims["is_admin"] != true {
		echo.Echo(echo.OrangeFG, fmt.Sprintf("PublicProfile patch handler error: expert_id %d does not match authenticated expert_id %d and 'is_admin' is %t", updated_public_profile.Expert_id, authenicaded_claims["id"], authenicaded_claims["is_admin"]))
		response.WriteHeader(403)
		return
	}

	err = repository.UpdateExpertPublicProfile(request.Context(), updated_public_profile)
	if err != nil {
		echo.Echo(echo.OrangeFG, "PublicProfile patch handler error: failed to update public_profile")
		echo.EchoErr(err)
		response.WriteHeader(500)
		return
	}

	echo.Echo(echo.GreenFG, "PublicProfile patch handler: public_profile updated successfully")
	response.WriteHeader(200)
	return
}

func postPublicProfileHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func getPublicProfileHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("id") != "" {
		getPublicProfileByIDHandler(response, request)
		return
	}

	response.WriteHeader(400)
	return
}

// CALLED BY: getPublicProfileHandler
func getPublicProfileByIDHandler(response http.ResponseWriter, request *http.Request) {
	var public_profile_id int

	public_profile_id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		echo.Echo(echo.OrangeFG, fmt.Sprintf("PublicProfile get by id handler error: failed to convert id '%s' to int", request.URL.Query().Get("id")))
		echo.EchoErr(err)
		response.WriteHeader(400)
		return
	}

	public_profile, err := repository.GetExpertPublicProfile(request.Context(), public_profile_id)
	if err != nil {
		echo.EchoErr(err)
		response.WriteHeader(404)
		return
	}

	if !public_profile.IsCreated() {
		public_profile.Create()
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(public_profile)

	return
}

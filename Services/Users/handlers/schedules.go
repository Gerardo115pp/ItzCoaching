package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"itz_customers_service/models"
	"itz_customers_service/repository"
	"itz_customers_service/server"
	"net/http"
	"strconv"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

func SchedulesHandler(order_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getSchedulesHandler(response, request)
		case http.MethodPost:
			postSchedulesHandler(response, request)
		case http.MethodPatch:
			patchSchedulesHandler(response, request)
		case http.MethodDelete:
			deleteSchedulesHandler(response, request)
		case http.MethodPut:
			putSchedulesHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func putSchedulesHandler(response http.ResponseWriter, request *http.Request) {
	var authenticated_claims map[string]interface{} = request.Context().Value("claim_data").(map[string]interface{})
	json_body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		echo.Echo(echo.RedFG, "Error reading request body")
		echo.EchoErr(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var schedule *models.ExpertSchedule = new(models.ExpertSchedule)

	err = json.Unmarshal(json_body, schedule)
	if err != nil {
		echo.Echo(echo.RedFG, "Error parsing request body")
		echo.EchoErr(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	if schedule.Expert_id != authenticated_claims["id"] {
		echo.Echo(echo.RedFG, "Unauthorized request")
		response.WriteHeader(http.StatusUnauthorized)
		return
	}

	if authenticated_claims["is_available"].(bool) != schedule.IsLogicallyAvailable() && !schedule.IsLogicallyAvailable() {
		echo.Echo(echo.PurpleFG, fmt.Sprintf("Updating expert %d availability, since she has no available days ", schedule.Expert_id))

		var expert *models.Expert
		expert, err = repository.GetExpertByID(request.Context(), schedule.Expert_id)
		if err != nil {
			echo.Echo(echo.RedFG, "Error getting expert")
			echo.EchoErr(err)
			response.WriteHeader(http.StatusInternalServerError)
			return
		}

		expert.IsAvailable = false

		err = repository.UpdateExpert(request.Context(), expert)
		if err != nil {
			echo.Echo(echo.RedFG, "Error updating expert")
			echo.EchoErr(err)
			response.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	public_profile, err := repository.GetExpertPublicProfile(request.Context(), schedule.Expert_id)
	if err != nil {
		echo.Echo(echo.RedFG, "Error getting expert profile")
		echo.EchoErr(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = public_profile.UpdateSchedule(schedule)
	if err != nil {
		echo.Echo(echo.RedFG, "Error updating expert schedule")
		echo.EchoErr(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	echo.Echo(echo.GreenFG, fmt.Sprintf("Expert with id %d updated her schedule", schedule.Expert_id))
	response.WriteHeader(200)
	return
}

func deleteSchedulesHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func patchSchedulesHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func postSchedulesHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func getSchedulesHandler(response http.ResponseWriter, request *http.Request) {
	var expert_id string = request.URL.Query().Get("id")
	if expert_id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	echo.Echo(echo.GreenFG, fmt.Sprintf("Getting expert schedule for expert with id: %s", expert_id))

	parsed_id, err := strconv.Atoi(expert_id)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error parsing expert id: %s", err.Error()))
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var public_profile *models.PublicProfile
	public_profile, err = repository.GetExpertPublicProfile(request.Context(), parsed_id)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error getting expert public profile for if %d", parsed_id))
		echo.EchoErr(err)
		response.WriteHeader(404)
		return
	}

	expert, err := repository.GetExpertByID(request.Context(), parsed_id)
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error getting expert with id %d", parsed_id))
		echo.EchoErr(err)
		response.WriteHeader(404)
		return
	}

	var schedule *models.ExpertSchedule
	schedule, err = public_profile.GetSchedule(expert.IsAvailable)

	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error getting expert schedule for expert with id: %d", parsed_id))
		echo.EchoErr(err)
		response.WriteHeader(500)
		return
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(schedule)
	return
}

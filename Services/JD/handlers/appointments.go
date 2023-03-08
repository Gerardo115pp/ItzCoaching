package handlers

import (
	"encoding/json"
	"fmt"
	"itz_JD_service/models"
	"itz_JD_service/repository"
	"itz_JD_service/server"
	"net/http"
	"strconv"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

type PostIsAppointmentAvaliableRequest struct {
	ExpertID int    `json:"expert_id"`
	UtcStart string `json:"utc_start"`
	UtcEnd   string `json:"utc_end"`
}

func AppointmentsHandler(order_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getAppointmentsHandler(response, request)
		case http.MethodPost:
			postAppointmentsHandler(response, request)
		case http.MethodPatch:
			patchAppointmentsHandler(response, request)
		case http.MethodDelete:
			deleteAppointmentsHandler(response, request)
		case http.MethodPut:
			putAppointmentsHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func putAppointmentsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func deleteAppointmentsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func patchAppointmentsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func postAppointmentsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func getAppointmentsHandler(response http.ResponseWriter, request *http.Request) {
	var request_path string = request.URL.Path

	switch request_path {
	case "/appointments/active":
		getActiveAppointmentsHandler(response, request)
	case "/appointments/customer":
		getAppointmentsByCustomerEmailHandler(response, request)
	case "/appointments/appointment":
		getAppointmentByIDHandler(response, request)
	case "/appointments/expert":
		getAppointmentsByExpertIDHandler(response, request)
	case "/appointments":
		getAppointments(response, request)
	default:
		response.WriteHeader(404)
	}
}

func getAppointmentsByCustomerEmailHandler(response http.ResponseWriter, request *http.Request) {
	var email string = request.URL.Query().Get("email")
	if email == "" {
		echo.Echo(echo.RedFG, "email is empty")
		response.WriteHeader(400)
		return
	}

	customer_appointments, err := repository.GetAllAppointmentsByCustomerEmail(email)
	if err != nil {
		echo.Echo(echo.RedFG, "error while getting appointments")
		echo.EchoErr(err)
		response.WriteHeader(500)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(customer_appointments)
}

func getActiveAppointmentsHandler(response http.ResponseWriter, request *http.Request) {
	active_appointments, err := repository.GetActiveAppointments()
	if err != nil {
		echo.Echo(echo.RedFG, "error while getting appointments")
		echo.EchoErr(err)
		response.WriteHeader(500)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(active_appointments)
}

func getAppointmentByIDHandler(response http.ResponseWriter, request *http.Request) {
	var appointment_id int

	appointment_id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		echo.Echo(echo.RedFG, "error while converting expert id")
		echo.EchoErr(err)
		response.WriteHeader(400)
		return
	}

	appointment, err := repository.GetAppointmentByID(appointment_id)
	if err != nil {
		echo.Echo(echo.RedFG, "error while getting appointments")
		echo.EchoErr(err)
		response.WriteHeader(404)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(appointment)
}

func postIsAppointmentAvaliableHandler(response http.ResponseWriter, request *http.Request) {

	var is_appointment_avaliable_request *PostIsAppointmentAvaliableRequest = new(PostIsAppointmentAvaliableRequest)

	err := json.NewDecoder(request.Body).Decode(is_appointment_avaliable_request)
	if err != nil {
		echo.Echo(echo.RedFG, "error while decoding request")
		echo.EchoErr(err)
		response.WriteHeader(400)
		return
	}

	is_avaliable, err := repository.IsTimeSlotAvailable(is_appointment_avaliable_request.ExpertID, is_appointment_avaliable_request.UtcStart, is_appointment_avaliable_request.UtcEnd)
	if err != nil {
		echo.Echo(echo.RedFG, "error while getting appointments")
		echo.EchoErr(err)
		response.WriteHeader(500)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	response.Write([]byte(fmt.Sprintf(`{"avaliable": %t}`, is_avaliable)))
}

func getAppointments(response http.ResponseWriter, request *http.Request) {
	var appointments []models.Appointment

	appointments, err := repository.GetAllAppointments()
	if err != nil {
		echo.Echo(echo.RedFG, "error while getting appointments")
		echo.EchoErr(err)
		response.WriteHeader(500)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(appointments)
}

func getAppointmentsByExpertIDHandler(response http.ResponseWriter, request *http.Request) {
	var expert_id int

	expert_id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		echo.Echo(echo.RedFG, "error while converting expert id")
		echo.EchoErr(err)
		response.WriteHeader(400)
		return
	}

	expert_appointments, err := repository.GetActiveAppointmentsByExpertID(expert_id)
	if err != nil {
		echo.Echo(echo.RedFG, "error while getting appointments")
		echo.EchoErr(err)
		response.WriteHeader(500)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(expert_appointments)
}

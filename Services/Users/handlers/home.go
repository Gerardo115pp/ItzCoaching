package handlers

import (
	"encoding/json"
	"net/http"

	"itz_customers_service/server"
)

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

package handlers

import (
	"itz_JD_service/server"
	"net/http"
)

func SystemStatsHandler(order_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getSystemStatsHandler(response, request)
		case http.MethodPost:
			postSystemStatsHandler(response, request)
		case http.MethodPatch:
			patchSystemStatsHandler(response, request)
		case http.MethodDelete:
			deleteSystemStatsHandler(response, request)
		case http.MethodPut:
			putSystemStatsHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(http.StatusOK)
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func putSystemStatsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func deleteSystemStatsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func patchSystemStatsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func postSystemStatsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func getSystemStatsHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusMethodNotAllowed)
	return
}

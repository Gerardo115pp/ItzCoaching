package handlers

import (
	"bonhart_auth_service/server"
	"net/http"
)

type CreateRecoverySessionRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	NewPassword   string `json:"new_password"`
	RecoveryToken string `json:"recovery_token"`
}

func RecoveryHandler(auth_server server.Server) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			getRecoveryHandler(response, request)
		case http.MethodPost:
			postRecoveryHandler(response, request)
		case http.MethodPatch:
			patchRecoveryHandler(response, request)
		case http.MethodDelete:
			deleteRecoveryHandler(response, request)
		case http.MethodPut:
			putRecoveryHandler(response, request)
		case http.MethodOptions:
			response.WriteHeader(200) // V8 crys really hard and patheticly when you don't return a status code
		default:
			response.WriteHeader(405)
		}
	}
}

func getRecoveryHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Get Requests */
	response.WriteHeader(405) // Method not allowed
}

// Send password recovery email
func postRecoveryHandler(response http.ResponseWriter, request *http.Request) {
	// 	/* Handle Post Requests */
	// 	var recovery_request CreateRecoverySessionRequest

	// 	err := json.NewDecoder(request.Body).Decode(&recovery_request)
	// 	if err != nil {
	// 		response.WriteHeader(400)
	// 		return
	// 	}

	// 	_, err = repository.GetCustomerByEmail(request.Context(), recovery_request.Email)
	// 	if err != nil {
	// 		echo.EchoWarn(fmt.Sprintf("Customer with email: %s not found(%s)", recovery_request.Email, err.Error()))
	// 		response.WriteHeader(200) // We don't want to give away if the email exists or not
	// 		return
	// 	}

	// 	recovery_session, err := repository.GetRecoverySessionByEmail(request.Context(), recovery_request.Email)
	// 	if err == nil {
	// 		// A password recovery was requested less than 24 hours ago
	// 		echo.Echo(echo.BlueBG, fmt.Sprintf("A password recovery was requested less than 24 hours ago for email: %s", recovery_request.Email))
	// 		response.WriteHeader(409) // Conflict, client must wait 24 hours before requesting a new recovery
	// 		return
	// 	}

	// 	recovery_session = models.CreateRecoverySession(recovery_request.Email)

	// 	err = repository.SetRecoverySession(request.Context(), recovery_session)
	// 	if err != nil {
	// 		echo.EchoWarn(fmt.Sprintf("Error setting recovery session: %s", err.Error()))
	// 		response.WriteHeader(500)
	// 		return
	// 	}

	// 	recovery_token, err := recovery_session.CreateRecoverySessionToken(app_config.JWT_SECRET)
	// 	recovery_url := fmt.Sprintf("%s/#/password-reset?recovery_token=%s", app_config.CUSTOMER_PAYMENT_APP, recovery_token)

	// 	send_email_func := func() {
	// 		err = app_email.SendRecoveryEmail(recovery_request.Email, recovery_url)
	// 		if err != nil {
	// 			echo.EchoErr(err)
	// 		}
	// 	}

	// 	go send_email_func()
	// 	// _ = send_email_func
	// 	// this is for testing purposes, uncomment the line above to send the email. keep in mind that there is a rate limit for sending emails
	// 	// echo.EchoDebug(fmt.Sprintf("Recovery URL: %s", recovery_url))

	// echo.Echo(echo.GreenBG, fmt.Sprintf("Recovery email sent to: %s", recovery_request.Email))
}

// Change password
func patchRecoveryHandler(response http.ResponseWriter, request *http.Request) {
	// /* Handle Patch Requests */
	// var reset_password_request ResetPasswordRequest

	// err := json.NewDecoder(request.Body).Decode(&reset_password_request)
	// if err != nil {
	// 	response.WriteHeader(400)
	// 	response.Write([]byte("Invalid request, missing fields"))
	// 	return
	// }

	// recovery_token, err := models.ParseRecoverySessionToken(app_config.JWT_SECRET, reset_password_request.RecoveryToken) // Parse and validate token
	// if err != nil {
	// 	echo.EchoErr(err)
	// 	response.WriteHeader(401)
	// 	return
	// }

	// recovery_session, err := repository.GetRecoverySessionByEmail(request.Context(), recovery_token.Email)
	// if err != nil {
	// 	response.WriteHeader(404)
	// 	return
	// }

	// if recovery_session.Used {
	// 	response.WriteHeader(403)
	// 	return
	// }

	// err = workflows.ChangeCustomerPassword(recovery_session.Email, reset_password_request.NewPassword)
	// if err != nil {
	// 	response.WriteHeader(500)
	// 	return
	// }

	// // customer must be deleted from local database so the new password can be used to login
	// err = repository.DeleteCustomerByEmail(request.Context(), recovery_session.Email)
	// if err != nil {
	// 	response.WriteHeader(500)
	// 	return
	// }

	// recovery_session.Used = true
	// err = repository.SetRecoverySession(request.Context(), recovery_session)

	// if err != nil {
	// 	response.WriteHeader(500)
	// 	return
	// }

	// echo.Echo(echo.GreenBG, fmt.Sprintf("Password reset for email: %s", recovery_token.Email))
}

func deleteRecoveryHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Delete Requests */
	response.WriteHeader(405) // Method not allowed
}

func putRecoveryHandler(response http.ResponseWriter, request *http.Request) {
	/* Handle Put Requests */
	response.WriteHeader(405) // Method not allowed
}

// Path: handlers/Recovery.go

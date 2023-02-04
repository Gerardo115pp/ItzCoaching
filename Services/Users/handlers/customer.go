package handlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	app_config "itz_customers_service/Config"
// 	"itz_customers_service/models"
// 	"itz_customers_service/repository"
// 	"itz_customers_service/server"
// 	"net/http"

// 	"github.com/Gerardo115pp/patriots_lib/echo"
// )

// /* REGISTER CUSTOMER */

// type SignupRequest struct {
// 	Username string `json:"username"`
// 	Name     string `json:"name"`
// 	Password string `json:"password"`
// 	Email    string `json:"email"`
// 	Whatsapp string `json:"whatsapp"`
// }

// type UpdateCustomerRequest struct {
// 	ID string `json:"id"`
// 	SignupRequest
// }

// func CustomerHandler(orders_server server.Server) http.HandlerFunc {
// 	return func(response http.ResponseWriter, request *http.Request) {
// 		switch request.Method {
// 		case http.MethodGet:
// 			getCustomerHandler(response, request)
// 		case http.MethodPost:
// 			postCustomerHandler(response, request)
// 		case http.MethodPatch:
// 			patchCustomerHandler(response, request)
// 		case http.MethodOptions:
// 			response.WriteHeader(http.StatusOK) // V8 crys really hard and patheticly when you don't return a status code
// 		default:
// 			response.WriteHeader(http.StatusMethodNotAllowed)
// 		}

// 	}
// }

// func getCustomerHandler(response http.ResponseWriter, request *http.Request) {
// 	var target_user *models.Customer
// 	var err error

// 	if customer_id := request.URL.Query().Get("id"); customer_id != "" {
// 		target_user, err = repository.GetCustomerByID(request.Context(), customer_id)
// 	} else if customer_email := request.URL.Query().Get("email"); customer_email != "" {
// 		target_user, err = repository.GetCustomerByEmail(request.Context(), customer_email)
// 	} else {
// 		response.WriteHeader(400)
// 		return
// 	}

// 	if err != nil {
// 		response.WriteHeader(404)
// 		echo.EchoErr(err)
// 		return
// 	}

// 	response.WriteHeader(http.StatusOK)
// 	response.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(response).Encode(target_user)
// }

// func postCustomerHandler(response http.ResponseWriter, request *http.Request) {
// 	var signup_request SignupRequest
// 	err := json.NewDecoder(request.Body).Decode(&signup_request)
// 	if err != nil {
// 		echo.EchoErr(err)
// 		response.WriteHeader(400)
// 		return
// 	}

// 	customer := models.CreateCustomerObject(signup_request.Name, signup_request.Email, signup_request.Password, signup_request.Whatsapp, 0)
// 	customer.Username = signup_request.Username
// 	err = repository.InsertCustomer(request.Context(), customer)
// 	if err != nil {
// 		echo.EchoErr(err)
// 		response.WriteHeader(500)
// 		return
// 	}

// 	response.WriteHeader(201)
// }

// func patchCustomerHandler(response http.ResponseWriter, request *http.Request) {
// 	// Change customer info. This request must come from an internal service, so we first check that the Authorization header == app_config.DOMAIN_SECRET

// 	var authentication_header string = request.Header.Get("Authorization")

// 	if authentication_header != app_config.DOMAIN_SECRET {
// 		echo.EchoErr(fmt.Errorf("Unauthorized access to PATCH /customer"))
// 		response.WriteHeader(401)
// 		return
// 	}

// 	var attribute_path string = request.URL.Path
// 	customer_update_request := new(UpdateCustomerRequest) // uses the same struct as the signup request, but fields can be empty

// 	err := json.NewDecoder(request.Body).Decode(&customer_update_request)
// 	if err != nil {
// 		echo.EchoErr(err)
// 		response.WriteHeader(400)
// 		return
// 	}

// 	// id must be provided or email
// 	if customer_update_request.ID == "" && customer_update_request.Email == "" {
// 		echo.Echo(echo.RedFG, "no Customer ID or Email provided")
// 		response.WriteHeader(400)
// 		return
// 	}

// 	echo.Echo(echo.PurpleFG, fmt.Sprintf("Updating attrubute %s", attribute_path))

// 	switch attribute_path {
// 	case "/customers/password":
// 		echo.Echo(echo.PurpleFG, "Validating password")
// 		if customer_update_request.Password == "" || customer_update_request.Email == "" {
// 			echo.Echo(echo.RedFG, "no password or email provided")
// 			response.WriteHeader(400)
// 			return
// 		}

// 		echo.Echo(echo.PurpleFG, fmt.Sprintf("Updating password for customer %s", customer_update_request.Email))

// 		target_customer, err := repository.GetCustomerByEmail(request.Context(), customer_update_request.Email)
// 		if err != nil {
// 			echo.EchoErr(err)
// 			response.WriteHeader(404)
// 			return
// 		}
// 		echo.Echo(echo.PurpleFG, fmt.Sprintf("Updating password for customer %s", target_customer.Email))

// 		target_customer.UpdatePassword(customer_update_request.Password)
// 		err = repository.UpdateCustomer(request.Context(), target_customer)
// 		if err != nil {
// 			echo.EchoErr(err)
// 			response.WriteHeader(500)
// 			return
// 		}

// 		echo.Echo(echo.GreenFG, "Success")
// 		response.WriteHeader(200)
// 		return
// 	default:
// 		response.WriteHeader(501)
// 		return
// 	}
// }

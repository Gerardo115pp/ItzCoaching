package main

import (
	"context"
	"fmt"
	app_config "itz_customers_service/Config"
	"itz_customers_service/database"
	"itz_customers_service/handlers"
	"itz_customers_service/middleware"
	"itz_customers_service/repository"
	"itz_customers_service/server"

	"github.com/Gerardo115pp/patriot_router"
	"github.com/Gerardo115pp/patriots_lib/echo"
)

func BinderRoutes(server server.Server, router *patriot_router.Router) {
	router.RegisterRoute(patriot_router.NewRoute("/", true), handlers.HomeHandler(server))
	router.RegisterRoute(patriot_router.NewRoute("/experts", true), middleware.CheckAuth(handlers.ExpertsHandler(server)))
	router.RegisterRoute(patriot_router.NewRoute("/public_profiles", true), middleware.CheckAuth(handlers.PublicProfilesHandler(server)))
	router.RegisterRoute(patriot_router.NewRoute("^/profile_pictures.*", false), middleware.CheckAuth(handlers.ProfilePicturesHandler(server)))
}

func main() {
	app_config.VerifyConfig()

	echo.Echo(echo.GreenFG, "Starting weiter_customers_service")

	var new_server_config *server.ServerConfig = new(server.ServerConfig)
	new_server_config.Port = app_config.CUST_PORT

	users_repository, err := database.NewMYSQLrepository()
	if err != nil {
		echo.Echo(echo.RedFG, "Error while creating users repository")
		echo.EchoFatal(err)
	}

	repository.SetUsersRepository(users_repository)

	echo.EchoDebug(fmt.Sprintf("server config: %+v", new_server_config))

	server, err := server.NewBroker(context.Background(), new_server_config)
	if err != nil {
		echo.EchoFatal(err)
	}

	server.Run(BinderRoutes)
}

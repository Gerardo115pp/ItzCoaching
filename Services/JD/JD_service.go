package main

import (
	"context"
	"fmt"
	app_config "itz_JD_service/Config"
	"itz_JD_service/database"
	"itz_JD_service/handlers"
	"itz_JD_service/middleware"
	"itz_JD_service/repository"
	"itz_JD_service/server"

	"github.com/Gerardo115pp/patriot_router"
	"github.com/Gerardo115pp/patriots_lib/echo"
)

func BinderRoutes(server server.Server, router *patriot_router.Router) {
	router.RegisterRoute(patriot_router.NewRoute("/", true), handlers.HomeHandler(server))
	router.RegisterRoute(patriot_router.NewRoute("/stats", true), middleware.CheckAuth(handlers.SystemStatsHandler(server)))
	router.RegisterRoute(patriot_router.NewRoute("^/appointments", false), middleware.CheckAuth(handlers.AppointmentsHandler(server)))
}

func main() {
	app_config.Init()

	echo.Echo(echo.GreenFG, "Starting weiter_customers_service")

	var new_server_config *server.ServerConfig = new(server.ServerConfig)
	new_server_config.Port = app_config.JD_PORT

	appointment_database, err := database.NewAppointmentRepository()
	if err != nil {
		echo.EchoFatal(err)
	}

	repository.SetAppointmentsRepository(appointment_database)

	echo.EchoDebug(fmt.Sprintf("server config: %+v", new_server_config))

	server, err := server.NewBroker(context.Background(), new_server_config)
	if err != nil {
		echo.EchoFatal(err)
	}

	server.Run(BinderRoutes)
}

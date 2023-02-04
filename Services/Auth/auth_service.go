package main

import (
	app_config "bonhart_auth_service/Config"
	"bonhart_auth_service/database"
	"bonhart_auth_service/handlers"
	"bonhart_auth_service/repository"
	"bonhart_auth_service/server"
	"context"
	"fmt"

	"github.com/Gerardo115pp/patriot_router"
	"github.com/Gerardo115pp/patriots_lib/echo"
)

func BinderRoutes(server server.Server, router *patriot_router.Router) {

	router.RegisterRoute(patriot_router.NewRoute("/", true), handlers.HomeHandler(server))
	router.RegisterRoute(patriot_router.NewRoute("/tokens/experts", false), handlers.ExpertsHandler(server))
	router.RegisterRoute(patriot_router.NewRoute("/tokens/admins", false), handlers.AdminsHandler(server))
	// router.RegisterRoute(patriot_router.NewRoute("/recovery", false), handlers.RecoveryHandler(server))
}

func main() {
	app_config.VerifyConfig()

	echo.EchoDebug(app_config.StrConfig())

	echo.Echo(echo.GreenFG, "Starting bonhart_auth_service")

	var new_server_config *server.ServerConfig = new(server.ServerConfig)
	new_server_config.Port = app_config.AUTH_PORT

	expert_repository, err := database.NewMYSQLrepository()
	if err != nil {
		echo.Echo(echo.RedFG, "Failed to connect to auth database")
		echo.EchoFatal(err)
	}

	// TODO: enable recovery sessions
	// recovery_sessions_repository, err := database.NewRedisRecoverySessionRepository()
	// if err != nil {
	// 	echo.Echo(echo.RedFG, "Failed to connect to redis")
	// 	echo.EchoFatal(err)
	// }

	repository.SetExpertRepository(expert_repository)

	echo.EchoDebug(fmt.Sprintf("server config: %+v", new_server_config))

	server, err := server.NewBroker(context.Background(), new_server_config)
	if err != nil {
		echo.EchoFatal(err)
	}

	server.Run(BinderRoutes)

}

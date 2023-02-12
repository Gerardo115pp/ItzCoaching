package app_config

import (
	"fmt"
	"os"

	"github.com/Gerardo115pp/patriots_lib/echo"
) // Loads the configuration from the environment variables
var JD_PORT string = os.Getenv("JD_PORT")
var JWT_SECRET string = os.Getenv("JWT_SECRET")
var DOMAIN_SECRET string = os.Getenv("DOMAIN_SECRET")
var MYSQL_USER string = os.Getenv("MYSQL_USER")
var MYSQL_PASSWORD string = os.Getenv("MYSQL_PASSWORD")
var MYSQL_HOST string = os.Getenv("MYSQL_HOST")
var MYSQL_DB string = os.Getenv("MYSQL_DB")
var MYSQL_PORT string = os.Getenv("MYSQL_PORT")
var OPERATION_DATA_PATH string = os.Getenv("OPERATION_DATA_PATH")

func verifyConfig() {
	if JD_PORT == "" {
		panic("JD_PORT environment variable is required")
	}
	if JWT_SECRET == "" {
		panic("JWT_SECRET environment variable is required")
	}
	if DOMAIN_SECRET == "" {
		panic("DOMAIN_SECRET environment variable is required")
	}
	if MYSQL_USER == "" {
		panic("MYSQL_USER environment variable is required")
	}
	if MYSQL_PASSWORD == "" {
		panic("MYSQL_PASSWORD environment variable is required")
	}
	if MYSQL_HOST == "" {
		panic("MYSQL_HOST environment variable is required")
	}
	if MYSQL_DB == "" {
		panic("MYSQL_DB environment variable is required")
	}
	if MYSQL_PORT == "" {
		panic("MYSQL_PORT environment variable is required")
	}
}

func Init() {
	verifyConfig()
	err := CreateConfig()
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error while creating config: %s", err.Error()))
		panic(err)
	}

	echo.Echo(echo.GreenFG, "Config loaded successfully")
}

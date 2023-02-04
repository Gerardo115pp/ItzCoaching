package app_config

import (
	"fmt"
	"os"
) // Loads the configuration from the environment variables

var MYSQL_USER string = os.Getenv("MYSQL_USER")
var MYSQL_PASSWORD string = os.Getenv("MYSQL_PASSWORD")
var MYSQL_HOST string = os.Getenv("MYSQL_HOST")
var MYSQL_DB string = os.Getenv("MYSQL_DB")
var MYSQL_PORT string = os.Getenv("MYSQL_PORT")
var AUTH_PORT string = os.Getenv("AUTH_PORT")
var JWT_SECRET string = os.Getenv("JWT_SECRET")
var DOMAIN_SECRET string = os.Getenv("DOMAIN_SECRET")
var REDIS_URL string = os.Getenv("REDIS_URL")
var REDIS_PASSWORD string = os.Getenv("REDIS_PASSWORD")
var MAIL_USERNAME string = os.Getenv("MAIL_USERNAME")
var MAIL_PASSWORD string = os.Getenv("MAIL_PASSWORD")
var MAIL_SERVER string = os.Getenv("MAIL_SERVER")
var MAIL_PORT string = os.Getenv("MAIL_PORT")
var MAIL_FROM string = os.Getenv("MAIL_FROM")

func VerifyConfig() {
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
	if AUTH_PORT == "" {
		panic("AUTH_PORT environment variable is required")
	}
	if JWT_SECRET == "" {
		panic("JWT_SECRET environment variable is required")
	}
	if DOMAIN_SECRET == "" {
		panic("DOMAIN_SECRET environment variable is required")
	}
	if REDIS_URL == "" {
		panic("REDIS_URL environment variable is required")
	}
	if MAIL_USERNAME == "" {
		panic("MAIL_USERNAME environment variable is required")
	}
	if MAIL_PASSWORD == "" {
		panic("MAIL_PASSWORD environment variable is required")
	}
	if MAIL_SERVER == "" {
		panic("MAIL_SERVER environment variable is required")
	}
	if MAIL_PORT == "" {
		panic("MAIL_PORT environment variable is required")
	}
	if MAIL_FROM == "" {
		panic("MAIL_FROM environment variable is required")
	}
}

func StrConfig() string {
	var config_str string = fmt.Sprintf("MYSQL_USER: %s", MYSQL_USER)
	config_str += fmt.Sprintf("MYSQL_PASSWORD: %s", MAIL_PASSWORD)
	config_str += fmt.Sprintf("\nMYSQL_HOST: %s", MYSQL_HOST)
	config_str += fmt.Sprintf("\nMYSQL_DB: %s", MYSQL_DB)
	config_str += fmt.Sprintf("\nMYSQL_PORT: %s", MYSQL_PORT)
	config_str += fmt.Sprintf("\nAUTH_PORT: %s", AUTH_PORT)
	config_str += fmt.Sprintf("\nJWT_SECRET: %s", JWT_SECRET)
	config_str += fmt.Sprintf("\nDOMAIN_SECRET: %s", DOMAIN_SECRET)
	config_str += fmt.Sprintf("\nREDIS_URL: %s", REDIS_URL)
	config_str += fmt.Sprintf("\nMAIL_USERNAME: %s", MAIL_USERNAME)
	config_str += fmt.Sprintf("\nMAIL_PASSWORD: %s", MAIL_PASSWORD)
	config_str += fmt.Sprintf("\nMAIL_SERVER: %s", MAIL_SERVER)
	config_str += fmt.Sprintf("\nMAIL_PORT: %s", MAIL_PORT)
	config_str += fmt.Sprintf("\nMAIL_FROM: %s", MAIL_FROM)

	return config_str
}

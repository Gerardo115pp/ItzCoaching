package app_config

import (
	"encoding/json"
	"fmt"
	"itz_JD_service/helpers"
	"os"

	"github.com/Gerardo115pp/patriots_lib/echo"
)

func CreateConfig() error {
	if !helpers.PathExists(OPERATION_DATA_PATH) {
		return fmt.Errorf("Operation data path not found: %s", OPERATION_DATA_PATH)
	}

	config_directory, err := createConfigDirectory()
	if err != nil {
		echo.Echo(echo.RedFG, fmt.Sprintf("Error while creating config directory: %s", err.Error()))
		echo.EchoFatal(err)
	}

	err = createStripeConfig(config_directory)

	return err
}

func createConfigDirectory() (string, error) {
	var err error
	var config_directory string = fmt.Sprintf("%s/system_config", OPERATION_DATA_PATH)

	if helpers.PathExists(config_directory) {
		echo.Echo(echo.GreenFG, fmt.Sprintf("Config directory already exists: %s", config_directory))
		return config_directory, nil
	}

	err = os.Mkdir(config_directory, os.ModePerm)
	if err != nil {
		return "", err
	}

	return config_directory, nil
}

func createStripeConfig(config_dir string) error {
	var stripe_config map[string]any = make(map[string]any)
	var stripe_config_path string = fmt.Sprintf("%s/stripe_config.json", config_dir)

	// CREATING STRIPE CONFIG
	echo.Echo(echo.PurpleFG, "Creating stripe config")
	stripe_config["secret_key"] = ""
	stripe_config["publishable_key"] = ""
	stripe_config["webhook_secret"] = ""
	stripe_config["webhook_url"] = ""
	stripe_config["webhook_port"] = ""
	stripe_config["webhook_enabled"] = false

	json_data, err := json.Marshal(stripe_config)
	if err != nil {
		return err
	}

	err = helpers.WriteFile(stripe_config_path, json_data)
	if err != nil {
		return err
	}

	return nil
}

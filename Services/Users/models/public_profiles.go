package models

import (
	"fmt"
	app_config "itz_customers_service/Config"
	"itz_customers_service/helpers"
	"os"
)

type PublicProfile struct {
	Public_name        string `json:"public_name"`
	Professional_title string `json:"professional_title"`
	Description        string `json:"description"`
	Biref              string `json:"biref"`
	Image_url          string `json:"image_url"`
	Expert_id          int    `json:"expert_id"`
}

func (pp *PublicProfile) IsCreated() bool {
	if pp.Expert_id == 0 {
		return false
	}

	var public_profile_dir string = fmt.Sprintf("%s/expert_%d", app_config.EXPERTS_DATA_PATH, pp.Expert_id)

	if _, err := os.Stat(public_profile_dir); os.IsNotExist(err) {
		return false
	}

	return true
}

func (pp *PublicProfile) Create() error {
	var public_profile_dir string = fmt.Sprintf("%s/expert_%d", app_config.EXPERTS_DATA_PATH, pp.Expert_id)
	err := os.MkdirAll(public_profile_dir, os.ModePerm)
	if err != nil {
		return err
	}

	var schedule_basic_template string = fmt.Sprintf("%s/schedule_basic_template.json", app_config.OPERATION_DATA_PATH)
	return helpers.CopyFile(schedule_basic_template, fmt.Sprintf("%s/schedule.json", public_profile_dir))
}

func (pp *PublicProfile) Update(new_data map[string]any) error {
	if new_data["public_name"] != nil {
		pp.Public_name = new_data["public_name"].(string)
	}

	if new_data["professional_title"] != nil {
		pp.Professional_title = new_data["professional_title"].(string)
	}

	if new_data["description"] != nil {
		pp.Description = new_data["description"].(string)
	}

	if new_data["biref"] != nil {
		pp.Biref = new_data["biref"].(string)
	}

	if new_data["image_url"] != nil {
		pp.Image_url = new_data["image_url"].(string)
	}

	return nil
}

package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	app_config "itz_customers_service/Config"
	"itz_customers_service/helpers"
	"os"
)

type PublicProfile struct {
	Public_name        string `json:"public_name"`
	Professional_title string `json:"professional_title"`
	Description        string `json:"description"`
	Brief              string `json:"brief"`
	Image_url          string `json:"image_url"`
	Expert_id          int    `json:"expert_id"`
	ExpertType         string `json:"expert_type,omitempty"`
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
	var public_profile_dir string = pp.getProfileDir()
	err := os.MkdirAll(public_profile_dir, os.ModePerm)
	if err != nil {
		return err
	}

	var schedule_basic_template string = fmt.Sprintf("%s/experts/schedules/schedule_basic_template.json", app_config.OPERATION_DATA_PATH)
	return helpers.CopyFile(schedule_basic_template, fmt.Sprintf("%s/schedule.json", public_profile_dir))
}

func (pp *PublicProfile) DestroyContent() error {
	var public_profile_dir string = fmt.Sprintf("%s/expert_%d", app_config.EXPERTS_DATA_PATH, pp.Expert_id)
	return os.RemoveAll(public_profile_dir)
}

func (pp PublicProfile) getProfileDir() string {
	return fmt.Sprintf("%s/expert_%d", app_config.EXPERTS_DATA_PATH, pp.Expert_id)
}

func (pp PublicProfile) GetSchedule(is_available bool) (*ExpertSchedule, error) {
	var schedule_path string = fmt.Sprintf("%s/schedule.json", pp.getProfileDir())
	schedule, err := parseSchedule(schedule_path)
	if err != nil {
		return nil, err
	}

	schedule.IsExpertAvailable = is_available
	schedule.Expert_id = pp.Expert_id

	return schedule, nil
}

func (pp *PublicProfile) UpdateSchedule(new_schedule *ExpertSchedule) error {
	new_schedule.Expert_id = pp.Expert_id
	json_data, err := json.Marshal(new_schedule)
	if err != nil {
		return err
	}

	schedule_path := fmt.Sprintf("%s/schedule.json", pp.getProfileDir())
	err = ioutil.WriteFile(schedule_path, json_data, 0644)
	return err
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
		pp.Brief = new_data["biref"].(string)
	}

	if new_data["image_url"] != nil {
		pp.Image_url = new_data["image_url"].(string)
	}

	return nil
}

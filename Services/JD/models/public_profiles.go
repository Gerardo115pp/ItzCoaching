package models

type PublicProfile struct {
	Public_name        string `json:"public_name"`
	Professional_title string `json:"professional_title"`
	Description        string `json:"description"`
	Brief              string `json:"brief"`
	Image_url          string `json:"image_url"`
	Expert_id          int    `json:"expert_id"`
	ExpertType         string `json:"expert_type,omitempty"`
}

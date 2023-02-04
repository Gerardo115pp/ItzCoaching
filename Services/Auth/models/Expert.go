package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Expert struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsActive    bool   `json:"is_active"`
	IsAvailable bool   `json:"is_available"`
	ExpertType  string `json:"expert"`
}

type ExpertTokenInfo struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	IsActive    bool   `json:"is_active"`
	IsAvailable bool   `json:"is_available"`
	ExpertType  string `json:"expert"`
}

func (expert *Expert) getTokenInfo() *ExpertTokenInfo {
	return &ExpertTokenInfo{
		ID:          expert.ID,
		Username:    expert.Username,
		Email:       expert.Email,
		IsActive:    expert.IsActive,
		IsAvailable: expert.IsAvailable,
		ExpertType:  expert.ExpertType,
	}
}

func CreateExpertObject(expert_data map[string]any) (*Expert, error) {
	var new_expert *Expert = new(Expert)

	if expert_data["id"] == nil {
		expert_data["id"] = 0
	}
	new_expert.ID = expert_data["id"].(int)

	if expert_data["username"] == nil {
		return nil, fmt.Errorf("Username is required")
	}
	new_expert.Username = expert_data["username"].(string)

	if expert_data["email"] == nil {
		return nil, fmt.Errorf("Email is required")
	}
	new_expert.Email = expert_data["email"].(string)

	if expert_data["password"] == nil {
		return nil, fmt.Errorf("Password is required")
	}

	raw_password := expert_data["password"].(string)
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(raw_password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	new_expert.Password = string(hashed_password)

	if expert_data["is_active"] == nil {
		expert_data["is_active"] = true
	}
	new_expert.IsActive = expert_data["is_active"].(bool)

	if expert_data["is_available"] == nil {
		expert_data["is_available"] = false
	}
	new_expert.IsAvailable = expert_data["is_available"].(bool)

	if expert_data["expert_type"] == nil {
		return nil, fmt.Errorf("Expert type is required")
	}
	new_expert.ExpertType = expert_data["expert_type"].(string)

	return new_expert, nil
}

func (expert Expert) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(expert.Password), []byte(password))
}

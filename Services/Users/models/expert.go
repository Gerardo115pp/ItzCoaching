package models

import (
	"fmt"
	"time"

	"github.com/Gerardo115pp/patriots_lib/echo"
	"golang.org/x/crypto/bcrypt"
)

type Expert struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	IsActive    bool      `json:"is_active"`
	IsAvailable bool      `json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
	LastLogin   time.Time `json:"last_login"`
	ExpertType  string    `json:"expert"`
	CreatedBy   int       `json:"created_by"`
}

func CreateExpertObject(expert_data map[string]any) (*Expert, error) {
	var new_expert *Expert = new(Expert)

	if expert_data["id"] == nil {
		expert_data["id"] = 0
	}
	new_expert.ID = expert_data["id"].(int)

	if expert_data["name"] == nil {
		return nil, fmt.Errorf("Name is required")
	}
	new_expert.Name = expert_data["name"].(string)

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

	if expert_data["created_at"] == nil {
		expert_data["created_at"] = time.Now()
	}

	new_expert.CreatedAt = expert_data["created_at"].(time.Time)

	if expert_data["last_login"] == nil {
		expert_data["last_login"] = time.Now()
	}

	new_expert.LastLogin = expert_data["last_login"].(time.Time)

	if expert_data["expert_type"] == nil {
		return nil, fmt.Errorf("Expert type is required")
	}
	new_expert.ExpertType = expert_data["expert_type"].(string)

	if expert_data["created_by"] == nil {
		return nil, fmt.Errorf("the admin id is required")
	}

	new_expert.CreatedBy = expert_data["created_by"].(int)

	return new_expert, nil
}

func (expert *Expert) PublicData() map[string]any {
	return map[string]any{
		"id":           expert.ID,
		"name":         expert.Name,
		"username":     expert.Username,
		"email":        expert.Email,
		"is_active":    expert.IsActive,
		"is_available": expert.IsAvailable,
		"created_at":   expert.CreatedAt,
		"last_login":   expert.LastLogin,
		"expert_type":  expert.ExpertType,
	}
}

func (expert *Expert) Update(new_data map[string]any) {
	if new_data["name"] != nil && new_data["name"] != "" {
		expert.Name = new_data["name"].(string)
	}
	if new_data["username"] != nil && new_data["username"] != "" {
		expert.Username = new_data["username"].(string)
	}
	if new_data["email"] != nil && new_data["email"] != "" {
		expert.Email = new_data["email"].(string)
	}
	if new_data["is_active"] != nil {
		expert.IsActive = new_data["is_active"].(bool)
	}
	if new_data["is_available"] != nil {
		expert.IsAvailable = new_data["is_available"].(bool)
	}
	if new_data["expert_type"] != nil && new_data["expert_type"] != "" {
		expert.ExpertType = new_data["expert_type"].(string)
	}
	if new_data["password"] != nil && new_data["password"] != "" {
		var new_password string = new_data["password"].(string)
		expert.UpdatePassword(new_password)
	}

}

func (expert *Expert) UpdatePassword(new_password string) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.DefaultCost)
	if err != nil {
		echo.EchoFatal(err)
	}
	expert.Password = string(hashed_password)
}

package models

import "golang.org/x/crypto/bcrypt"

type ItzAdmin struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	IsActive     bool   `json:"is_active"`
	IsSuperadmin bool   `json:"is_superadmin"`
}

type AdminTokenInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

func (admin *ItzAdmin) getTokenInfo() *AdminTokenInfo {
	return &AdminTokenInfo{
		ID:       admin.ID,
		Username: admin.Username,
		Email:    admin.Email,
		IsActive: admin.IsActive,
	}
}

func (expert ItzAdmin) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(expert.Password), []byte(password))
}

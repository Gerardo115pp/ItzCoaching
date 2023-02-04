package repository

import (
	"context"
	"itz_customers_service/models"
)

type UsersRepository interface {
	InsertExpert(ctx context.Context, expert *models.Expert) error
	GetExpertByID(ctx context.Context, id int) (*models.Expert, error)
	GetExpertByEmail(ctx context.Context, email string) (*models.Expert, error)
	UpdateExpert(ctx context.Context, expert *models.Expert) error
	DeleteExpert(ctx context.Context, id int) error
	GetExpertPublicProfile(ctx context.Context, expert_id int) (*models.PublicProfile, error)
	UpdateExpertPublicProfile(ctx context.Context, public_profile *models.PublicProfile) error
	Close() error
}

var user_repo_implentation UsersRepository

func SetUsersRepository(repository UsersRepository) {
	user_repo_implentation = repository
}

func InsertExpert(ctx context.Context, expert *models.Expert) error {
	return user_repo_implentation.InsertExpert(ctx, expert)
}

func GetExpertByID(ctx context.Context, id int) (*models.Expert, error) {
	return user_repo_implentation.GetExpertByID(ctx, id)
}

func GetExpertByEmail(ctx context.Context, email string) (*models.Expert, error) {
	return user_repo_implentation.GetExpertByEmail(ctx, email)
}

func UpdateExpert(ctx context.Context, expert *models.Expert) error {
	return user_repo_implentation.UpdateExpert(ctx, expert)
}

func DeleteExpert(ctx context.Context, id int) error {
	return user_repo_implentation.DeleteExpert(ctx, id)
}

func GetExpertPublicProfile(ctx context.Context, expert_id int) (*models.PublicProfile, error) {
	return user_repo_implentation.GetExpertPublicProfile(ctx, expert_id)
}

func UpdateExpertPublicProfile(ctx context.Context, public_profile *models.PublicProfile) error {
	return user_repo_implentation.UpdateExpertPublicProfile(ctx, public_profile)
}

func Close() error {
	return user_repo_implentation.Close()
}

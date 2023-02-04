package repository

import (
	"bonhart_auth_service/models"
	"context"
)

type CustomerRepository interface {
	GetAdminByID(ctx context.Context, id int) (*models.ItzAdmin, error)
	GetAdminByEmail(ctx context.Context, email string) (*models.ItzAdmin, error)
	GetExpertByID(ctx context.Context, id int) (*models.Expert, error)
	GetExpertByEmail(ctx context.Context, email string) (*models.Expert, error)
	Close() error
}

var auth_repository CustomerRepository

func SetExpertRepository(repository CustomerRepository) {
	auth_repository = repository
}

func GetExpertByID(ctx context.Context, id int) (*models.Expert, error) {
	return auth_repository.GetExpertByID(ctx, id)
}

func GetExpertByEmail(ctx context.Context, email string) (*models.Expert, error) {
	return auth_repository.GetExpertByEmail(ctx, email)
}

func GetAdminByID(ctx context.Context, id int) (*models.ItzAdmin, error) {
	return auth_repository.GetAdminByID(ctx, id)
}

func GetAdminByEmail(ctx context.Context, email string) (*models.ItzAdmin, error) {
	return auth_repository.GetAdminByEmail(ctx, email)
}

func closeRepository() error {
	return auth_repository.Close()
}

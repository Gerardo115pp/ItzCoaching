package repository

import (
	"bonhart_auth_service/models"
	"context"
)

type RecoverySessionRepository interface {
	SetRecoverySession(ctx context.Context, recoverySession *models.RecoverySession) error
	GetRecoverySessionByEmail(ctx context.Context, email string) (*models.RecoverySession, error)
	Close() error
}

var recovery_session_repo_implentation RecoverySessionRepository

func SetRecoverySessionRepository(repository RecoverySessionRepository) {
	recovery_session_repo_implentation = repository
}

func SetRecoverySession(ctx context.Context, recovery_session *models.RecoverySession) error {
	return recovery_session_repo_implentation.SetRecoverySession(ctx, recovery_session)
}

func GetRecoverySessionByEmail(ctx context.Context, email string) (*models.RecoverySession, error) {
	return recovery_session_repo_implentation.GetRecoverySessionByEmail(ctx, email)
}

func closeRecoverySessionRepo() error {
	return recovery_session_repo_implentation.Close()
}

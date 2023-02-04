package database

import (
	app_config "bonhart_auth_service/Config"
	"bonhart_auth_service/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Gerardo115pp/patriots_lib/echo"
	"github.com/go-redis/redis"
)

type RedisRecoverySessionRepository struct {
	client *redis.Client
}

func NewRedisRecoverySessionRepository() (*RedisRecoverySessionRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     app_config.REDIS_URL,
		Password: app_config.REDIS_PASSWORD,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	echo.Echo(echo.GreenFG, "Redis connection established", pong)

	return &RedisRecoverySessionRepository{client: client}, nil
}

func (rrs *RedisRecoverySessionRepository) SetRecoverySession(ctx context.Context, recovery_session *models.RecoverySession) error {
	var rrs_key = fmt.Sprintf("recovery-session-%s", recovery_session.Email)
	var key_ttl time.Duration = time.Hour * 24

	/* Check if this key exists and if so use the same ttl */
	old_ttl, err := rrs.client.TTL(rrs_key).Result()
	if err == nil {
		if old_ttl != -2*time.Second {
			key_ttl = old_ttl
		}
	} else {
		echo.EchoWarn(fmt.Sprintf("Error getting TTL for key %s: %s", rrs_key, err.Error()))
	}

	recovery_session_serialized, err := json.Marshal(recovery_session)
	if err != nil {
		return err
	}

	err = rrs.client.Set(rrs_key, string(recovery_session_serialized), key_ttl).Err()
	return err
}

func (rrs *RedisRecoverySessionRepository) GetRecoverySessionByEmail(ctx context.Context, email string) (*models.RecoverySession, error) {
	var rrs_key = fmt.Sprintf("recovery-session-%s", email)

	val, err := rrs.client.Get(rrs_key).Result()
	if err != nil {
		return nil, err
	}

	var recovery_session models.RecoverySession
	err = json.Unmarshal([]byte(val), &recovery_session)
	if err != nil {
		return nil, err
	}

	return &recovery_session, nil
}

func (rrs *RedisRecoverySessionRepository) Close() error {
	err := rrs.client.Close()
	if err != nil {
		echo.EchoWarn(fmt.Sprintf("Error closing redis connection: %s", err.Error()))
	}

	return err
}

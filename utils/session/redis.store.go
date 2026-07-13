package session

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionStore interface {
	save(ctx context.Context,sessionId string,data map[string]string,expiry time.Duration) error
	get(ctx context.Context,sessionId string) (map[string]string,error)
	destroy(ctx context.Context,sessionId string) error
}

type SessionManager struct{
	redis_client *redis.Client
}

func NewSessionStore(redis_client *redis.Client) *SessionManager{
	return &SessionManager{
		redis_client: redis_client,
	}
}

func( store *SessionManager) save(ctx context.Context,sessionId string,data map[string]string,expiry time.Duration) error{


	



}
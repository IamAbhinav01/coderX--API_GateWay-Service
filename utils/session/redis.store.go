package session

import (
	"Coderx/utils/json"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionStore interface {
	Save(ctx context.Context,sessionId string,data map[string]string,expiry time.Duration) error
	Get(ctx context.Context,sessionId string) (map[string]string,error)
	Destroy(ctx context.Context,sessionId string) error
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

	bytes,err := json.Marshall(data)

	if err != nil{
		fmt.Println("Error Happend while Marshalling data")
		return err
	}

	return store.redis_client.Set(ctx,"session:"+sessionId,bytes,expiry).Err()

}

func (store * SessionManager) Get(ctx context.Context,sessionId string) (map[string]string,error){

	bytes,err := store.redis_client.Get(ctx,"session:"+sessionId).Bytes()

	if err != nil{
		fmt.Println("Error occured while retrieved data from redis")
		return nil,err
	}

	var data map[string]string
	err = json.UnMarshall(bytes,&data)
	return data,err

}
package app

import (
	"Coderx/DB/repositories"
	"Coderx/config/db"
	"Coderx/config/env"
	"Coderx/controllers"
	"Coderx/router"
	"Coderx/services"
	"Coderx/utils/session"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type Application struct {
	Config string
}


func NewApplication() *Application {
	return &Application{
		Config: env.GetString("PORT"),
	}
}

func (app *Application) Run() error {

	DB, err := db.InitDB()

	if err != nil {
		fmt.Println("Error while initializing the database:", err)
		return err
	}

	fmt.Println("Database initialized successfully")

	var redisClient *redis.Client
	
	redisClient = redis.NewClient(&redis.Options{
		Addr: env.GetString("REDIS_ADDR"),
	})

	sessionStore := session.NewSessionStore(redisClient)
	sessionManager := session.NewSessionManager(sessionStore)

	user_repository := repositories.NewUserRepository(DB)
	user_service := services.NewService(user_repository)
	user_controller := controllers.NewController(user_service,sessionManager)
	fmt.Println("User controller initialized")

	// Set up router with chi
	appRouter := router.SetUpRouter(user_controller)

	// Format port address
	addr := app.Config
	if !strings.HasPrefix(addr, ":") {
		addr = ":" + addr
	}

	server := http.Server{
		Addr:         addr,
		Handler:      appRouter,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server running on", server.Addr)
	return server.ListenAndServe()
}
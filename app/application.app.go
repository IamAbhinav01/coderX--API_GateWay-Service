package app

import (
	"Coderx/DB/repositories"
	"Coderx/config/db"
	"Coderx/config/env"
	"Coderx/controllers"
	"Coderx/router"
	"Coderx/services"
	"fmt"
	"net/http"
	"strings"
	"time"
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

	user_repository := repositories.NewUserRepository(DB)
	user_service := services.NewService(user_repository)
	user_controller := controllers.NewController(user_service)
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
package app

import (
	"Coderx/DB/repositories"
	"Coderx/config/db"
	"Coderx/config/env"
	"Coderx/services"
	"fmt"
	"net/http"
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

func (app *Application) Run() error{

	DB,err:=db.InitDB()

	if err != nil{
		fmt.Println("Error while initalising the database")
	}

	fmt.Println("Databse Initialised ",DB)

	user_repository := repositories.NewUserRepository(DB)
	user_service := services.NewService(user_repository)
	fmt.Println("user service : ",user_service)


	
	server := http.Server{
		Addr: app.Config,
		Handler: http.NewServeMux(),
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}
	fmt.Println("Server running on PORT",server.Addr)
	return server.ListenAndServe()
}
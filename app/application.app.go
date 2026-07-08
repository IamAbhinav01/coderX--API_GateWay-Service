package app

import (
	"Coderx/config/db"
	"Coderx/config/env"
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

	server := http.Server{
		Addr: app.Config,
		Handler: http.NewServeMux(),
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}
	fmt.Println("Server running on PORT",server.Addr)
	return server.ListenAndServe()
}
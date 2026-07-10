package router

import (
	"Coderx/controllers"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router){

	r.Post("/SignUp",controllers.S)


}
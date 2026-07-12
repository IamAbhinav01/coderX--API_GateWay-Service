package router

import (
	"Coderx/controllers"
	"Coderx/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, userController *controllers.UserController) {

	r.With(middlewares.SingUpRequestValidation).Post("/SignUp", userController.SignUp)

}
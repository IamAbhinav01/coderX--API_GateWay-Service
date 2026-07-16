package router

import (
	"Coderx/controllers"
	"Coderx/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, userController *controllers.UserController) {

	r.With(middlewares.SignUpRequestValidation).Post("/SignUp", userController.SignUp)
	r.Post("/Login",userController.Login)

}
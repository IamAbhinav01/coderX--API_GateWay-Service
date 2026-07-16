package router

import (
	"Coderx/controllers"
	"Coderx/middlewares"
	"Coderx/utils/session"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, userController *controllers.UserController, sm *session.SessionManager) {

	r.With(middlewares.SignUpRequestValidation).Post("/SignUp", userController.SignUp)
	r.With(middlewares.LoginRequestValidation).Post("/Login",userController.Login)
	r.With(middlewares.AuthMiddleware(sm)).Get("/greet",userController.Greetings)
	r.With(middlewares.AuthMiddleware(sm)).Post("/LogOut",userController.LogOut)
}
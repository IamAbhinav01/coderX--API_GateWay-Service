package router

import (
	"Coderx/controllers"
	"Coderx/utils/session"

	"github.com/go-chi/chi/v5"
)

func SetUpRouter(userController *controllers.UserController,sm *session.SessionManager) *chi.Mux {

	router := chi.NewRouter()

	router.Route("/api/v1", func(r chi.Router) {
		RegisterUserRoutes(r, userController,sm)
	})

	return router
}
package router

import "github.com/go-chi/chi/v5"

func SetUpRouter() *chi.Mux{

	router := chi.NewRouter()

	router.Route("/api/v1",func(r chi.Router) {
		
	})
}
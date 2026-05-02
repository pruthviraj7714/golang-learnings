package routes

import (
	"orm/handlers"

	"github.com/go-chi/chi/v5"
)

func RegsiterRouter(h *handlers.UserHandler, r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.GetAllUsers)
		r.Get("/{id}", h.GetUserById)
		r.Post("/", h.CreateUser)
		r.Delete("/{id}", h.DeleteUser)
		r.Put("/{id}", h.UpdateUser)
	})
}

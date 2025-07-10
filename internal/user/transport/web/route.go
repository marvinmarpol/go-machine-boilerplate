package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Route interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

func RegisterRoutes(r chi.Router, route Route) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HELLOO")
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/", route.Create)
		r.Get("/{id}", route.Get)
	})

}

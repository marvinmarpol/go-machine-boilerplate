package web

import (
	"database/sql"
	"go-machine-boilerplate/internal/user/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Handler struct {
	service *service.UserService
}

func NewHandler(r *chi.Mux, service *service.UserService) *Handler {
	handler := &Handler{service: service}
	RegisterRoutes(r, handler)

	return handler
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserRequest

	if err := render.DecodeJSON(r.Body, &payload); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{Message: err.Error()})
		return
	}

	result, err := h.service.Create(payload.Email, payload.Name)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, ErrorResponse{Message: err.Error()})
		return
	}

	render.JSON(w, r, result)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.Get(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		if err == sql.ErrNoRows {
			render.Status(r, http.StatusNotFound)
		}
		render.JSON(w, r, ErrorResponse{Message: err.Error()})
		return
	}

	render.JSON(w, r, result)
}

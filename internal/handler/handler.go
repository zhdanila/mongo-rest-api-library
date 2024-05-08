package handler

import (
	"mongo/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func(h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /get", h.get)
	mux.HandleFunc("GET /get/{name}", h.getByName)
	mux.HandleFunc("POST /create", h.create)
	mux.HandleFunc("PUT /update/{id}", h.update)
	mux.HandleFunc("DELETE /delete/{id}", h.delete)

	return mux
}
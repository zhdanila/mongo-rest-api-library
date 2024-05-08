package handler

import (
	"encoding/json"
	"fmt"
	"mongo/internal/models"
	"net/http"
)

func(h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Create(book)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("book was created - %s", id)))
}

func(h *Handler) get(w http.ResponseWriter, r *http.Request) {

}

func(h *Handler) getById(w http.ResponseWriter, r *http.Request) {

}

func(h *Handler) update(w http.ResponseWriter, r *http.Request) {

}

func(h *Handler) delete(w http.ResponseWriter, r *http.Request) {

}
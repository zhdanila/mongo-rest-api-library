package handler

import (
	"encoding/json"
	"fmt"
	"mongo/internal/models"
	"net/http"
	"strings"
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
	books, err := h.services.GetAll()
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	marshalledBooks, err := json.Marshal(&books)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledBooks)
}

func (h *Handler) getByName(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		NewErrorResponse(w, http.StatusBadRequest, "name parameter is missing")
		return
	}
	name := parts[2]

	book, err := h.services.GetById(name)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	marshalledBook, err := json.Marshal(&book)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledBook)
}


func(h *Handler) update(w http.ResponseWriter, r *http.Request) {
	var updatedBook models.Book

	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Update(updatedBook)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("book was created - %s", id)))
}

func(h *Handler) delete(w http.ResponseWriter, r *http.Request) {

}
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rflorezeam/libro-delete/services"
)

type Handler struct {
	service services.LibroService
}

func NewHandler(service services.LibroService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) EliminarLibro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID no proporcionado"})
		return
	}

	err := h.service.EliminarLibro(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Libro eliminado exitosamente"})
} 
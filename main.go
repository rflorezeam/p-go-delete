package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rflorezeam/libro-delete/config"
	"github.com/rflorezeam/libro-delete/repositories"
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

	result, err := h.service.EliminarLibro(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Libro no encontrado"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"mensaje": "Libro eliminado correctamente",
		"result":  result,
	})
}

func main() {
	// Inicializar la base de datos
	config.ConectarDB()

	// Inicializar las capas
	repo := repositories.NewLibroRepository()
	service := services.NewLibroService(repo)
	handler := NewHandler(service)
	
	// Configurar el router
	router := mux.NewRouter()
	router.HandleFunc("/libros/{id}", handler.EliminarLibro).Methods("DELETE")

	// Configurar el puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}

	fmt.Printf("Servicio de eliminación de libros corriendo en puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
} 
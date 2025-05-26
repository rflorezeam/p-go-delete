package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rflorezeam/libro-delete/config"
	"github.com/rflorezeam/libro-delete/handlers"
	"github.com/rflorezeam/libro-delete/repositories"
	"github.com/rflorezeam/libro-delete/services"
)

func main() {
	// Inicializar la base de datos
	config.ConectarDB()

	// Inicializar las capas
	repo := repositories.NewLibroRepository()
	service := services.NewLibroService(repo)
	handler := handlers.NewHandler(service)
	
	// Configurar el router
	router := mux.NewRouter()
	router.HandleFunc("/libros/{id}", handler.EliminarLibro).Methods("DELETE")

	// Configurar el puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}

	fmt.Printf("Servicio de eliminación de libros corriendo en puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
} 
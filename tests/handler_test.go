package tests

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rflorezeam/libro-delete/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLibroService struct {
	mock.Mock
}

func (m *MockLibroService) EliminarLibro(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestEliminarLibro_Exitoso(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	libroID := "123"
	mockService.On("EliminarLibro", libroID).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/libros/"+libroID, nil)
	w := httptest.NewRecorder()

	// Configurar las variables de ruta
	vars := map[string]string{
		"id": libroID,
	}
	req = mux.SetURLVars(req, vars)

	// Act
	handler.EliminarLibro(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Libro eliminado exitosamente", response["mensaje"])
	mockService.AssertExpectations(t)
}

func TestEliminarLibro_Error(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	libroID := "123"
	expectedError := errors.New("error al eliminar libro")
	mockService.On("EliminarLibro", libroID).Return(expectedError)

	req := httptest.NewRequest(http.MethodDelete, "/libros/"+libroID, nil)
	w := httptest.NewRecorder()

	// Configurar las variables de ruta
	vars := map[string]string{
		"id": libroID,
	}
	req = mux.SetURLVars(req, vars)

	// Act
	handler.EliminarLibro(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	
	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedError.Error(), response["error"])
	mockService.AssertExpectations(t)
}

func TestEliminarLibro_IDVacio(t *testing.T) {
	// Arrange
	mockService := new(MockLibroService)
	handler := handlers.NewHandler(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/libros/", nil)
	w := httptest.NewRecorder()

	// Configurar las variables de ruta vacías
	vars := map[string]string{}
	req = mux.SetURLVars(req, vars)

	// Act
	handler.EliminarLibro(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "ID no proporcionado", response["error"])
} 
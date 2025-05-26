package tests

import (
	"errors"
	"testing"

	"github.com/rflorezeam/libro-delete/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLibroRepository struct {
	mock.Mock
}

func (m *MockLibroRepository) EliminarLibro(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestEliminarLibro_ServicioExitoso(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)
	service := services.NewLibroService(mockRepo)

	libroID := "123"
	mockRepo.On("EliminarLibro", libroID).Return(nil)

	// Act
	err := service.EliminarLibro(libroID)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEliminarLibro_ServicioError(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)
	service := services.NewLibroService(mockRepo)

	libroID := "123"
	expectedError := errors.New("error al eliminar libro de la base de datos")
	mockRepo.On("EliminarLibro", libroID).Return(expectedError)

	// Act
	err := service.EliminarLibro(libroID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestEliminarLibro_ServicioIDVacio(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)
	service := services.NewLibroService(mockRepo)

	// Act
	err := service.EliminarLibro("")

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "ID no puede estar vacío", err.Error())
	mockRepo.AssertNotCalled(t, "EliminarLibro")
} 
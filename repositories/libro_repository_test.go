package repositories

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockLibroRepository struct {
	mock.Mock
}

func (m *MockLibroRepository) EliminarLibro(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestEliminarLibro_RepositorioExitoso(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)

	libroID := "123"
	mockRepo.On("EliminarLibro", libroID).Return(nil)

	// Act
	err := mockRepo.EliminarLibro(libroID)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEliminarLibro_RepositorioError(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)

	libroID := "123"
	expectedError := errors.New("error de conexión con la base de datos")
	mockRepo.On("EliminarLibro", libroID).Return(expectedError)

	// Act
	err := mockRepo.EliminarLibro(libroID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestEliminarLibro_RepositorioIDInvalido(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)

	libroID := "invalid-id"
	expectedError := errors.New("ID inválido")
	mockRepo.On("EliminarLibro", libroID).Return(expectedError)

	// Act
	err := mockRepo.EliminarLibro(libroID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestEliminarLibro_RepositorioLibroNoEncontrado(t *testing.T) {
	// Arrange
	mockRepo := new(MockLibroRepository)

	libroID := "123"
	expectedError := mongo.ErrNoDocuments
	mockRepo.On("EliminarLibro", libroID).Return(expectedError)

	// Act
	err := mockRepo.EliminarLibro(libroID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
} 
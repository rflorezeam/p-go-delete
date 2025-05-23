package services

import (
	"github.com/rflorezeam/libro-delete/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibroService interface {
	EliminarLibro(id string) (*mongo.DeleteResult, error)
}

type libroService struct {
	repo repositories.LibroRepository
}

func NewLibroService(repo repositories.LibroRepository) LibroService {
	return &libroService{
		repo: repo,
	}
}

func (s *libroService) EliminarLibro(id string) (*mongo.DeleteResult, error) {
	return s.repo.EliminarLibro(id)
} 
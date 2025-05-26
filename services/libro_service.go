package services

import (
	"errors"

	"github.com/rflorezeam/libro-delete/repositories"
)

type LibroService interface {
	EliminarLibro(id string) error
}

type libroService struct {
	repo repositories.LibroRepository
}

func NewLibroService(repo repositories.LibroRepository) LibroService {
	return &libroService{
		repo: repo,
	}
}

func (s *libroService) EliminarLibro(id string) error {
	if id == "" {
		return errors.New("ID no puede estar vacío")
	}

	return s.repo.EliminarLibro(id)
} 
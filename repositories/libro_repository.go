package repositories

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibroRepository interface {
	EliminarLibro(id string) error
}

type libroRepository struct {
	collection *mongo.Collection
}

func NewLibroRepository() LibroRepository {
	return &libroRepository{}
}

func (r *libroRepository) EliminarLibro(id string) error {
	if id == "" {
		return errors.New("ID no puede estar vacío")
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	result, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
} 
package repositories

import (
	"context"

	"github.com/rflorezeam/libro-delete/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibroRepository interface {
	EliminarLibro(id string) (*mongo.DeleteResult, error)
}

type libroRepository struct {
	collection *mongo.Collection
}

func NewLibroRepository() LibroRepository {
	return &libroRepository{
		collection: config.GetCollection(),
	}
}

func (r *libroRepository) EliminarLibro(id string) (*mongo.DeleteResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return r.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
} 
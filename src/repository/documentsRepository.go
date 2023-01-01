package repository

import (
	"context"

	"github.com/IngAngelMartinez/go-api-multitenant/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DocumentsRepository struct {
	Collection *mongo.Collection
}

var ctx = context.TODO()

func (d *DocumentsRepository) GetAll() ([]models.Document, error) {

	var documents []models.Document

	cur, err := d.Collection.Find(ctx, bson.D{})

	if err != nil {

		return nil, err
	}

	defer cur.Close(ctx)

	err = cur.All(ctx, &documents)

	if err != nil {

		return nil, err
	}

	return documents, nil
}

func (d *DocumentsRepository) GetById(id string) (models.Document, error) {

	var document models.Document

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		return document, err
	}

	err = d.Collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&document)

	if err == mongo.ErrNoDocuments {

		return document, nil
	}

	if err != nil {

		return document, err
	}

	return document, nil
}

func (d *DocumentsRepository) Create(document models.Document) (string, error) {

	var result string

	insert, err := d.Collection.InsertOne(ctx, document)

	if err != nil {

		return result, err
	}

	result = insert.InsertedID.(primitive.ObjectID).Hex()

	return result, nil
}

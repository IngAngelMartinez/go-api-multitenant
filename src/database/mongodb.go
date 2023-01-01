package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Uri          string
	DatabaseName string

	Client   *mongo.Client
	Database *mongo.Database
}

func (m *MongoDB) Connect() error {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))

	if err != nil {

		return err
	}

	m.Client = client

	m.Database = m.Client.Database(m.DatabaseName)

	return nil
}

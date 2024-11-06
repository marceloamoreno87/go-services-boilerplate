package database

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
}

func (d Mongo) InitDB(dsn string) (client *mongo.Client, err error) {
	client, err = mongo.Connect(options.Client().ApplyURI(dsn))
	if err != nil {
		return
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return
	}

	return
}

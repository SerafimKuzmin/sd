package flags

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBFlags struct {
	URL string `toml:"url"`
}

func (f MongoDBFlags) Init() (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(f.URL))

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	db, err := client.Database("recommend"), nil

	if err != nil {
		return nil, err
	}

	return db, nil
}

package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//InitDb represent a factory of database
func InitDb(ctx context.Context) (DBConn *mongo.Client) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	DBConn, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = DBConn.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return DBConn
}

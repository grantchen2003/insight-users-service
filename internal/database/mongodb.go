package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	client *mongo.Client
}

func (mongodb *MongoDb) Connect() error {
	mongodbUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongodbUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	mongodb.client = client

	log.Println("connected to MongoDB")

	return nil
}

func (mongodb *MongoDb) Close() error {
	if err := mongodb.client.Disconnect(context.TODO()); err != nil {
		return err
	}

	log.Println("connection to MongoDB closed")

	return nil
}

func (mongodb *MongoDb) SaveUser(sessionId string) (string, error) {
	user := struct{ SessionId string }{SessionId: sessionId}

	insertResult, err := mongodb.getCollection().InsertOne(context.Background(), user)

	if err != nil {
		return "", err
	}

	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (mongodb *MongoDb) GetUserBySessionId(sessionId string) (*User, error) {
	var result struct {
		Id        string `bson:"_id,omitempty"`
		SessionId string
	}

	filter := bson.D{{"sessionid", sessionId}}

	if err := mongodb.getCollection().FindOne(context.TODO(), filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &User{Id: result.Id, SessionId: result.SessionId}, nil
}

func (mongodb *MongoDb) getCollection() *mongo.Collection {
	databaseName := os.Getenv("MONGODB_DATABASE_NAME")
	collectionName := os.Getenv("MONGODB_COLLECTION_NAME")

	return mongodb.client.Database(databaseName).Collection(collectionName)
}

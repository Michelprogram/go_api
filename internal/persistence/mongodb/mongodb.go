package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client *mongo.Client
	ctx    context.Context
}

func (m Mongo) GetCollection(collection string) *mongo.Collection {
	return m.client.Database("Go_api").Collection(collection)
}

func NewMongo() (Mongo, error) {

	clientOptions := options.Client().ApplyURI("mongodb+srv://root:root@cluster0-fuans.mongodb.net/Crypto?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return Mongo{}, err
	}

	return Mongo{client: client, ctx: ctx}, nil
}

/*
	collection := client.Database("Go_api").Collection("Students")

	//	db, _ := client.ListDatabaseNames(ctx, bson.M{})

	//cursor, err := collection.Find(ctx, bson.M{})

	cursor, err := collection.Find(ctx, bson.M{"value": 42.29})

	if err != nil {
		return err
	}

	var crypto []bson.M
	if err = cursor.All(ctx, &crypto); err != nil {
		return err
	}

	fmt.Println(crypto)

	//	fmt.Println(db)
*/

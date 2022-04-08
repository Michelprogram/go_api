package persistence

import (
	"context"
	"fmt"
	"internal/entities"
	"internal/persistence/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentDaoMongoDB struct {
}

//var _ StudentDao = (*StudentDaoMongoDB)(nil)

var database mongodb.Mongo
var collection *mongo.Collection

func NewStudentDaoMongo() StudentDaoMongoDB {
	database, err := mongodb.NewMongo()

	if err != nil {
		fmt.Println("Error connexion")
	}

	collection = database.GetCollection("cryptomonnaie")

	return StudentDaoMongoDB{}
}

func (s StudentDaoMongoDB) FindAll() ([]entities.Student, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var crypto []bson.M
	cursor.All(ctx, &crypto)

	fmt.Println(crypto)

	collection.InsertOne(ctx, bson.D{
		{Key: "name", Value: "Test Students"},
	})

	return nil, err
}

func (s StudentDaoMongoDB) Find(id int) (*entities.Student, error) {
	return nil, nil
}

func (s StudentDaoMongoDB) Exists(id int) bool {
	return false
}

func (s StudentDaoMongoDB) Delete(id int) bool {
	return false
}

func (s StudentDaoMongoDB) Create(student entities.Student) bool {
	return false
}

func (s StudentDaoMongoDB) Update(student entities.Student) bool {
	return false
}

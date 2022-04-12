package persistence

import (
	"context"
	"fmt"
	"internal/entities"
	"internal/persistence/mongodb"
	"log"
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

	collection = database.GetCollection("Students")

	return StudentDaoMongoDB{}
}

func (s StudentDaoMongoDB) FindAll() []entities.Student {

	var students []entities.Student

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil
	}

	cursor.All(ctx, &students)

	return students

}

func (s StudentDaoMongoDB) Find(id int) (*entities.Student, error) {

	var student bson.M

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err := collection.FindOne(ctx, bson.D{{"id", id}}).Decode(&students)
	if err != nil {
		log.Fatal(err)
	}

	return *student, nil
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

/* 	collection.InsertOne(ctx, bson.D{
	{Key: "name", Value: "Test Students"},
})

	st := entities.Student{
	Id:             1,
	LastName:       "Dorian",
	FirstName:      "Gauron",
	Age:            21,
	LanguageDeCode: "FR",
}

rrr, err := collection.InsertOne(ctx, st)

if err != nil {
	fmt.Println(err)
}

fmt.Println(rrr.InsertedID)
*/

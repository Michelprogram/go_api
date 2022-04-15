package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"internal/entities"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type MyMongo struct {
	client *mongo.Client
}

func NewMyMongo() MyMongo {

	if client != nil {
		return MyMongo{client: client}
	}

	myMongo := openConnection()

	myMongo.insertFakeDataStudents()
	myMongo.insertFakeDataLanguages()

	return myMongo

}

func (m *MyMongo) insertFakeDataLanguages() {

	var collection string = "Languages"

	m.dropCollection(collection)

	var languages []entities.Language = []entities.Language{
		entities.NewLanguage(2, "FR", "France mongo"),
		entities.NewLanguage(1, "DE", "Allemagne mongo"),
		entities.NewLanguage(3, "CH", "Chine mongo"),
	}

	for _, language := range languages {

		res, _ := json.Marshal(language)

		m.Create(collection, string(res))
	}

}

func (m *MyMongo) insertFakeDataStudents() {

	var collection string = "Students"

	m.dropCollection(collection)

	var students []entities.Student = []entities.Student{
		entities.NewStudent(1, "Gaspar mongo", "Missiaen", 21, "FR"),
		entities.NewStudent(2, "Daurian mongo", "Gauron", 20, "DA"),
		entities.NewStudent(4, "Christopher mongo", "Lessirard", 20, "CH"),
		entities.NewStudent(3, "Daryl mongo", "Caruso", 20, "DE"),
	}

	for _, student := range students {

		res, _ := json.Marshal(student)

		m.Create(collection, string(res))

	}

}

func (m *MyMongo) dropCollection(collection string) {

	var coll *mongo.Collection = m.client.Database("Go_api").Collection(collection)

	err := coll.Drop(context.TODO())

	if err != nil {
		log.Fatal("Impossible de clean la collection students : ", err)
	}

}

func openConnection() MyMongo {

	var login string
	var password string
	var uri string

	err := godotenv.Load("./configs/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	login = os.Getenv("LOGIN_MONGODB")
	password = os.Getenv("PASSWORD_MONGODB")

	uri = fmt.Sprintf("mongodb+srv://%s:%s@cluster0-fuans.mongodb.net/Crypto?retryWrites=true&w=majority", login, password)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("Erreur de connexion à la base mongoDB")
	}

	fmt.Println("📁 Connecté à la base mongodb")

	return MyMongo{client: client}
}

func (m *MyMongo) GetAll(collection string) []string {

	var results []bson.M
	var data []string
	var coll *mongo.Collection = m.client.Database("Go_api").Collection(collection)

	cursor, err := coll.Find(context.Background(), bson.D{})

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatalf("Impossible de récupérer l'ensemble des données dans %s : %v \n", collection, err)
	}

	data = make([]string, len(results))

	for index, element := range results {
		delete(element, "_id")

		jsonStr, err := json.Marshal(element)

		if err != nil {
			log.Fatalf("Impossible de convertir un %s en json\n", collection)
		}

		data[index] = string(jsonStr)
	}

	return data

}

func (m *MyMongo) Get(collection string, key string, value interface{}) string {

	var result []bson.M
	var data string
	var coll *mongo.Collection = m.client.Database("Go_api").Collection(collection)

	var option *options.FindOptions = options.Find().SetProjection(bson.D{{"_id", 0}})

	cursor, err := coll.Find(context.Background(), bson.D{{key, value}}, option)

	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Fatalf("Impossible de récupérer l'ensemble des données dans %s : %v \n", collection, err)
	}

	if len(result) == 0 {
		return ""
	}

	dataf, err := json.Marshal(result[0])

	data = string(dataf)

	return data

}

func (m *MyMongo) Delete(collection string, key string, value interface{}) bool {

	var coll *mongo.Collection = m.client.Database("Go_api").Collection(collection)

	cursor, err := coll.DeleteOne(context.TODO(), bson.D{{key, value}})

	if err != nil {
		log.Fatalf("Impossible de supprimer la pair %s, %v pour la collection %s, err : %v\n", key, value, collection, err)
	}

	if cursor.DeletedCount != 0 {
		return true
	}

	return false

}

func (m *MyMongo) Create(collection string, data string) bool {

	var student entities.Student

	var coll *mongo.Collection = m.client.Database("Go_api").Collection(collection)

	json.Unmarshal([]byte(data), &student)

	_, err := coll.InsertOne(context.TODO(), student)

	if err != nil {
		log.Fatal("Problème lors de l'insertion")
		return false
	}

	return true
}

func (m *MyMongo) Update(collection string, data string) bool {

	var student entities.Student

	var coll *mongo.Collection = m.client.Database("Go_api").Collection(collection)

	json.Unmarshal([]byte(data), &student)

	_, err := coll.UpdateOne(context.TODO(), bson.D{{"id", student.Id}}, bson.D{{
		"$set", student,
	}})

	if err != nil {
		log.Fatal("Problème lors de l'insertion")
		return false
	}

	return true
}

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

	return openConnection()

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

	fmt.Println("Test : ", student)

	json.Unmarshal([]byte(data), &student)

	fmt.Println("Test : ", student)

	cursor, err := coll.InsertOne(context.TODO(), student)

	if err != nil {
		log.Fatal("Problème lors de l'insertion")
		return false
	}

	fmt.Println("Insertion de  : ", cursor)

	return true
}

func (m *MyMongo) Update(collection string, data string) {

}

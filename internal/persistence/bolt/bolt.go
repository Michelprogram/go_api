package bolt

import (
	"encoding/json"
	"fmt"
	"log"

	"internal/entities"

	"github.com/boltdb/bolt"
)

var database *bolt.DB

type MyBolt struct {
	db *bolt.DB
}

func GetMyBolt() MyBolt {

	if database != nil {
		return MyBolt{db: database}
	}

	return openMyBolt()

}

func openMyBolt() MyBolt {

	db, err := bolt.Open("myBolt.db", 0600, nil)

	database = db

	if err != nil {
		log.Fatal(err)
	}

	return MyBolt{db: database}
}

func (b *MyBolt) CreateDatabase() {

	var bucketsName []string = []string{"Students", "Languages"}

	for _, name := range bucketsName {
		b.deleteBucket(name)
		b.createBucket(name)
	}

	b.insertFakeDataStudents()
	b.insertFakeDataLanguages()
}

func (b *MyBolt) Close() {
	b.db.Close()
}

func (b *MyBolt) deleteBucket(bucketName string) {

	err := b.db.Update(func(tx *bolt.Tx) error {

		err := tx.DeleteBucket([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Le bucket %s ne peut être surpprimé car il n'éxiste pas.\n", bucketName)
	}
}

func (b *MyBolt) createBucket(bucketName string) {

	err := b.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (b *MyBolt) insertFakeDataStudents() {

	var students []entities.Student = []entities.Student{
		entities.NewStudent(1, "Gaspar bolt", "Missiaen", 21, "FR"),
		entities.NewStudent(2, "Daurian bolt", "Gauron", 20, "DA"),
		entities.NewStudent(4, "Christopher bolt", "Lessirard", 20, "CH"),
		entities.NewStudent(3, "Daryl bolt", "Caruso", 20, "DE"),
	}

	for _, student := range students {

		res, _ := json.Marshal(student)

		idStr := fmt.Sprintf("%d", student.Id)

		b.Put("Students", idStr, string(res))
	}

}

func (b *MyBolt) insertFakeDataLanguages() {

	var languages []entities.Language = []entities.Language{
		entities.NewLanguage(2, "FR", "France bolt"),
		entities.NewLanguage(1, "DE", "Allemagne bolt"),
		entities.NewLanguage(3, "CH", "Chine bolt"),
	}

	for _, language := range languages {

		res, _ := json.Marshal(language)

		b.Put("Languages", language.Code, string(res))
	}

}

func (b *MyBolt) Put(bucketName string, key string, value string) {

	err := b.db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(bucketName))

		if bucket == nil {
			panic("Bucket : " + bucketName + "existe pas")
		}

		bucket.Put([]byte(key), []byte(value))
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}

func (b *MyBolt) Get(bucketName string, key string) string {
	var value string

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		if bucket == nil {
			panic("Le Bucket : " + bucketName + " n'éxiste pas !!!")
		}

		value = string(bucket.Get([]byte(key)))
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return value
}

func (b *MyBolt) GetAll(bucketName string) []string {

	var resultat []string

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			value := fmt.Sprintf("%s", v)
			resultat = append(resultat, value)
			//fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return resultat
}

func (b *MyBolt) Delete(bucketName string, key string) error {

	err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		if bucket == nil {
			panic("Bucket : " + bucketName + "n'éxiste pas.")
		}

		bucket.Delete([]byte(key))

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

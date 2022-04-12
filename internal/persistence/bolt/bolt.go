package bolt

import (
	"encoding/json"
	"fmt"
	"log"

	"internal/entities"

	"github.com/boltdb/bolt"
)

type MyBolt struct {
	db *bolt.DB
}

func NewMyBolt() MyBolt {

	db, err := bolt.Open("myBolt.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	return MyBolt{db: db}
}

func (b *MyBolt) CreateDatabase() {

	b.createBucket("Students")
	//b.createBucket(db, "Language")
	b.insertFakeDataStudents()

}

func (b *MyBolt) Close() {
	b.db.Close()
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
		entities.NewStudent(1, "Gaspar", "Missiaen", 21, "23"),
		entities.NewStudent(2, "Daurian", "Gauron", 20, "Go"),
		entities.NewStudent(4, "Christopher", "Lessirard", 20, "26"),
		entities.NewStudent(3, "Daryl", "Caruso", 20, "-2"),
	}

	for _, student := range students {

		res, _ := json.Marshal(student)

		idStr := fmt.Sprintf("%d", student.Id)

		b.Put("Students", idStr, string(res))
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
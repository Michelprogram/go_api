package daomongo

import (
	"encoding/json"
	"errors"
	"internal/entities"
	"internal/persistence/interfaces"
	"internal/persistence/mongodb"
	"log"
)

type StudentDaoMongoDB struct {
}

var _ interfaces.StudentDao = (*StudentDaoMongoDB)(nil)

var myMongo mongodb.MyMongo = mongodb.NewMyMongo()

var collection = "Students"

func NewStudentDaoMongo() StudentDaoMongoDB {
	return StudentDaoMongoDB{}
}

func (s StudentDaoMongoDB) FindAll() []entities.Student {

	var students []entities.Student

	res := myMongo.GetAll(collection)

	students = make([]entities.Student, len(res))

	for index, student := range res {
		var st entities.Student

		json.Unmarshal([]byte(student), &st)

		students[index] = st
	}

	return students

}

func (s StudentDaoMongoDB) Find(id int) (*entities.Student, error) {

	var student entities.Student

	var res string = myMongo.Get(collection, "id", id)

	if res == "" {
		return nil, errors.New("L'id n'éxiste pas")
	}

	json.Unmarshal([]byte(res), &student)

	return &student, nil
}

func (s StudentDaoMongoDB) Exists(id int) bool {

	if myMongo.Get(collection, "id", id) != "" {
		return true
	}

	return false
}

func (s StudentDaoMongoDB) Delete(id int) bool {
	return myMongo.Delete(collection, "id", id)
}

func (s StudentDaoMongoDB) Create(student entities.Student) bool {

	studentStr, err := json.Marshal(student)

	if err != nil {
		log.Fatal("Problème lors de la conversion student to json byte")
		return false
	}

	if !s.Exists(student.Id) {
		return myMongo.Create(collection, string(studentStr))
	}

	return false

}

func (s StudentDaoMongoDB) Update(student entities.Student) bool {

	studentStr, err := json.Marshal(student)

	if err != nil {
		log.Fatal("Problème lors de la conversion student to json byte")
		return false
	}

	if s.Exists(student.Id){
		return myMongo.Update(collection, string(studentStr))
	}

	return false
}

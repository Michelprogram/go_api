package persistence

import (
	"encoding/json"
	"errors"
	"fmt"
	"internal/entities"
	"internal/persistence/bolt"
)

type StudentDaoBolt struct {
}

//var _ StudentDao = (*StudentDaoBolt)(nil)

var myBolt bolt.MyBolt = bolt.NewMyBolt()

func NewStudentDaoBolt() StudentDaoBolt {

	myBolt.CreateDatabase()

	return StudentDaoBolt{}
}

func (d *StudentDaoBolt) Find(id int) (*entities.Student, error) {

	var student entities.Student

	idStr := fmt.Sprintf("%d", id)

	res := myBolt.Get("Students", idStr)

	if res == "" {
		return nil, errors.New("L'id n'existe pas")
	}

	json.Unmarshal([]byte(res), &student)

	return &student, nil
}

func (d *StudentDaoBolt) Exists(id int) bool {

	idStr := fmt.Sprintf("%d", id)

	res := myBolt.Get("Students", idStr)

	if res == "" {
		return false
	}

	return true

}

func (d *StudentDaoBolt) Delete(id int) bool {

	idStr := fmt.Sprintf("%d", id)

	myBolt.Delete("Students", idStr)

	return true
}

/*
func FindAll() []entities.Student {

	return nil
}

/*






func Create(student entities.Student) bool {

}

func Update(student entities.Student) bool {

}
*/

package daobolt

import (
	"encoding/json"
	"errors"
	"fmt"
	"internal/entities"
	"internal/persistence/bolt"
	"internal/persistence/interfaces"
)

type StudentDaoBolt struct {
}

var _ interfaces.StudentDao = (*StudentDaoBolt)(nil)

var myBolt bolt.MyBolt = bolt.GetMyBolt()

func NewStudentDaoBolt() StudentDaoBolt {
	return StudentDaoBolt{}
}

func (d StudentDaoBolt) Find(id int) (*entities.Student, error) {

	var student entities.Student

	idStr := fmt.Sprintf("%d", id)

	res := myBolt.Get("Students", idStr)

	if res == "" {
		return nil, errors.New("L'id n'existe pas")
	}

	json.Unmarshal([]byte(res), &student)

	return &student, nil
}

func (d StudentDaoBolt) Exists(id int) bool {

	idStr := fmt.Sprintf("%d", id)

	res := myBolt.Get("Students", idStr)

	if res == "" {
		return false
	}

	return true

	//TODO : Revoir GET PB avec nombre qui commence par même chiffre 2 et insertion 21

}

func (d StudentDaoBolt) Delete(id int) bool {

	idStr := fmt.Sprintf("%d", id)

	err := myBolt.Delete("Students", idStr)

	if err != nil {
		return false
	}

	return true
}

func (d StudentDaoBolt) FindAll() []entities.Student {

	var students []entities.Student

	for _, student := range myBolt.GetAll("Students") {
		var st entities.Student
		json.Unmarshal([]byte(student), &st)
		students = append(students, st)
	}

	return students
}

func (d StudentDaoBolt) Create(student entities.Student) bool {

	res, _ := json.Marshal(student)

	idStr := fmt.Sprintf("%d", student.Id)

	if !d.Exists(student.Id) {
		myBolt.Put("Students", idStr, string(res))
		return true
	}

	return false

}

func (d StudentDaoBolt) Update(student entities.Student) bool {

	if d.Exists(student.Id) {

		res, _ := json.Marshal(student)

		idStr := fmt.Sprintf("%d", student.Id)

		myBolt.Delete("Students", idStr)

		myBolt.Put("Students", idStr, string(res))

		return true
	}

	return false
}

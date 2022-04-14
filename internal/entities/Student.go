package entities

import (
	"fmt"
)

type Student struct {
	Id             int    //`json:"Id"`
	FirstName      string //`json:"Field Str"`
	LastName       string //`json:"Field Str"`
	Age            int    //`json:"Field Int"`
	LanguageDeCode string //`json:"Field Str"`
}

//Check if Student implement de stringer
var _ fmt.Stringer = (*Student)(nil)

func NewStudent(id int, FirstName string, LastName string, Age int, LanguageDeCode string) Student {
	return Student{
		Id:             id,
		FirstName:      FirstName,
		LastName:       LastName,
		Age:            Age,
		LanguageDeCode: LanguageDeCode,
	}
}

func (s Student) String() string {
	return fmt.Sprintf("L'étudiant avec l'id %d, %s, %s, Age : %d, Dév : %s \n", s.Id, s.FirstName, s.LastName, s.Age, s.LanguageDeCode)
}

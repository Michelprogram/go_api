package entities

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Id             int                `bson:"id,omitempty"`
	FirstName      string             `bson:"firstname,omitempty"`
	LastName       string             `bson:"lastname,omitempty"`
	Age            int                `bson:"age,omitempty"`
	LanguageDeCode string             `bson:"languagedecode,omitempty"`
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

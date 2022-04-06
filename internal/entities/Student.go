package entities

import "fmt"

type Student struct {
	Id             int
	FirstName      string
	LastName       string
	Age            int
	LanguageDeCode string
}

func NewStudent(id int, FirstName string, LastName string, Age int, LanguageDeCode string) Student {
	return Student{id, FirstName, LastName, Age, LanguageDeCode}
}

func (s Student) String() string {
	return fmt.Sprintf("L'étudiant avec l'id %d, %s, %s, Age : %d, Dév : %s \n", s.Id, s.FirstName, s.LastName, s.Age, s.LanguageDeCode)
}

package persistence

import (
	"internal/entities"
	"internal/persistence"
)

var students []entities.Student = []entities.Student{
	entities.NewStudent(1, "Gaspar", "Missiaen", 21, "23"),
	entities.NewStudent(2, "Daurian", "Gauron", 20, "Go"),
	entities.NewStudent(4, "Christopher", "Lessirard", 20, "26"),
	entities.NewStudent(3, "Daryl", "Caruso", 20, "-2"),
}

type StudentDaoMemory struct {
}

var _ persistence.StudentDao = (*StudentDaoMemory)(nil)

func NewStudentDaoMemory() StudentDaoMemory {
	return StudentDaoMemory{}
}

func (s StudentDaoMemory) FindAll() []entities.Student {
	return students
}

func (s StudentDaoMemory) Find(id int) *entities.Student {

}

func (s StudentDaoMemory) Exists(id int) bool {

}

func (s StudentDaoMemory) Delete(id int) bool {

}

func (s StudentDaoMemory) Update(student entities.Student) bool {

}

func (s StudentDaoMemory) Create(student entities.Student) bool {

}

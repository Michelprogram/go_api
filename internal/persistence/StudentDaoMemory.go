package persistence

import (
	"internal/entities"
	"internal/persistence"
)

type StudentDaoMemory struct {
}

var _ persistence.StudentDao = (*StudentDaoMemory)(nil)

func NewStudentDaoMemory() StudentDaoMemory {
	return StudentDaoMemory{}
}

func (s StudentDaoMemory) FindAll() []entities.Student {

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

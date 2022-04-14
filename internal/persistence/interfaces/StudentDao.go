package interfaces

import "internal/entities"

type StudentDao interface {
	FindAll() []entities.Student
	Find(id int) (*entities.Student, error)
	Exists(id int) bool
	Delete(id int) bool
	Create(student entities.Student) bool
	Update(student entities.Student) bool
}

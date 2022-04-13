package daomemory

import (
	"errors"
	"internal/entities"
	ps "internal/persistence"
)

var students []entities.Student = []entities.Student{
	entities.NewStudent(1, "Gaspar", "Missiaen", 21, "23"),
	entities.NewStudent(2, "Daurian", "Gauron", 20, "Go"),
	entities.NewStudent(4, "Christopher", "Lessirard", 20, "26"),
	entities.NewStudent(3, "Daryl", "Caruso", 20, "-2"),
}

type StudentDaoMemory struct {
}

var _ ps.StudentDao = (*StudentDaoMemory)(nil)

func NewStudentDaoMemory() StudentDaoMemory {
	return StudentDaoMemory{}
}

func (s StudentDaoMemory) FindAll() []entities.Student {

	var newStudents []entities.Student = students

	for i := 0; i < len(newStudents)-1; i++ {
		j := i + 1

		stA := newStudents[i]
		stB := newStudents[j]

		if stA.Id > stB.Id {

			newStudents[i], newStudents[j] = stB, stA
		}
	}

	return newStudents
}

func (s StudentDaoMemory) Find(id int) (*entities.Student, error) {

	for _, student := range students {
		if student.Id == id {
			return &student, nil
		}
	}

	return nil, errors.New("L'id n'existe pas")
}

func (s StudentDaoMemory) Exists(id int) bool {

	for _, student := range students {
		if student.Id == id {
			return true
		}
	}

	return false
}

func (s StudentDaoMemory) Delete(id int) bool {

	for index, student := range students {
		if student.Id == id {
			students = append(students[:index], students[index+1:]...)
			return true
		}
	}

	return false

}

func (s StudentDaoMemory) Update(student entities.Student) bool {

	if !s.Exists(student.Id) {

		for index, element := range students {
			if student.Id == element.Id {
				students[index] = student
				return true
			}
		}
	}

	return false
}

func (s StudentDaoMemory) Create(student entities.Student) bool {

	if !s.Exists(student.Id) {
		students = append(students, student)
		return true
	}

	return false

}

package tests

import (
	"internal/entities"
	memomry "internal/persistence/daomemory"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dao memomry.StudentDaoMemory = memomry.NewStudentDaoMemory()

func TestFindAll(t *testing.T) {

	var res []int

	for _, student := range dao.FindAll() {
		res = append(res, student.Id)
	}

	assert.Equal(t, []int{1, 2, 3, 4}, res, "Students should be sort by ID")
}

func TestFindId1(t *testing.T) {

	var expected entities.Student = entities.NewStudent(1, "Gaspar memory", "Missiaen", 21, "FR")
	student, _ := dao.Find(1)

	assert.Equal(t, expected, *student, "Return student with id 1")

}

func TestFaillId10000(t *testing.T) {

	_, err := dao.Find(1000)

	if err == nil {
		t.Error("L'id 1000 ne doit pas exister dans la liste")
	}

}

func TestExist(t *testing.T) {
	assert.Equal(t, false, dao.Exists(1000), "L'id 1000 ne doit pas exister dans la liste des données")
	assert.Equal(t, true, dao.Exists(1), "L'id 1 doit exister dans la liste des données")
}

func TestDeleteId1(t *testing.T) {

	assert.Equal(t, false, dao.Delete(10000), "L'id 1000 n'existe pas.")
	assert.Equal(t, true, dao.Delete(1), "L'id 1 n'a pas été supprimé.")
	assert.Equal(t, false, dao.Exists(1), "L'id 1 n'a pas vraiment été supprimé")

}

func TestAddStudent(t *testing.T) {

	var student entities.Student = entities.NewStudent(21, "Gauron memory", "Dorian", 21, "FR")

	size := len(dao.FindAll())

	dao.Create(student)

	assert.Equal(t, size+1, len(dao.FindAll()), "L'étudiant n'a pas été ajouté")
}

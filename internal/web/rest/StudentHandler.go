package rest

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"io/ioutil"
	"net/http"

	"strconv"

	ps "internal/persistence"

	"github.com/gorilla/mux"
)

var students []entities.Student = []entities.Student{
	entities.NewStudent(1, "Gaspar", "Missiaen", 21, "23"),
	entities.NewStudent(2, "Daurian", "Gauron", 20, "Go"),
	entities.NewStudent(3, "Daryl", "Caruso", 20, "-2"),
	entities.NewStudent(4, "Christopher", "Lessirard", 20, "26"),
}

func StudentById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	dao := ps.NewStudentDaoMemory()

	student, err := dao.Find(id)

	if err == nil {
		res, _ := json.Marshal(*student)
		fmt.Fprintf(w, "%s", res)
		return
	}

	fmt.Fprintf(w, "L'étudiant avec l'id %d n'éxiste pas.", id)
}

func AllStudents(w http.ResponseWriter, r *http.Request) {

	var dao ps.StudentDaoMemory = ps.NewStudentDaoMemory()

	res, _ := json.Marshal(dao.FindAll())

	fmt.Fprintf(w, "%s", res)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var student entities.Student

	json.Unmarshal(reqBody, &student)

	students = append(students, student)

	res, _ := json.Marshal(student)

	fmt.Fprintf(w, "%s", res)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	for index, element := range students {
		if element.Id == id {

			res, _ := json.Marshal(element)

			students = append(students[:index], students[index+1:]...)

			fmt.Fprintf(w, "%s", res)
			return
		}
	}

	fmt.Fprintf(w, "L'étudiant avec l'id %d n'éxiste pas.", id)
}

func PutStudent(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var student entities.Student

	json.Unmarshal(reqBody, &student)

	for index, element := range students {
		if element.Id == student.Id {
			students[index] = student

			fmt.Fprintf(w, "%s", student)
			return
		}
	}

	fmt.Fprintf(w, "L'étudiant avec l'id %d n'éxiste pas.", student.Id)

}

//lgu.univ@gmail.com

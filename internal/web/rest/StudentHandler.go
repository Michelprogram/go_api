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

var dao ps.StudentDao = ps.NewStudentDaoMemory()

var daoMongodb ps.StudentDaoMongoDB = ps.NewStudentDaoMongo()

var daoBolt ps.StudentDaoBolt = ps.NewStudentDaoBolt()

func StudentById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	student, err := dao.Find(id)

	if err == nil {
		res, _ := json.Marshal(*student)
		fmt.Fprintf(w, "%s", res)
		return
	}

	fmt.Fprintf(w, "L'étudiant avec l'id %d n'éxiste pas.", id)
}

func AllStudents(w http.ResponseWriter, r *http.Request) {

	//res, _ := json.Marshal(dao.FindAll())

	res, _ := json.Marshal(daoMongodb.FindAll())

	//res, _ := json.Marshal(dao.FindAll())

	fmt.Fprintf(w, "%s", res)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var student entities.Student

	json.Unmarshal(reqBody, &student)

	if dao.Create(student) {
		res, _ := json.Marshal(student)

		fmt.Fprintf(w, "%s", res)

		return
	}

	fmt.Fprintf(w, "L'étudiant existe déjà")

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	student, err := dao.Find(id)

	if err != nil {
		fmt.Fprintf(w, "L'étudiant avec l'id %d, n'éxiste pas.", id)
		return
	}

	if dao.Delete(id) {
		res, _ := json.Marshal(student)
		fmt.Fprintf(w, "%s", res)
	}
}

func PutStudent(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var student entities.Student

	json.Unmarshal(reqBody, &student)

	if dao.Update(student) {
		res, _ := json.Marshal(student)
		fmt.Fprintf(w, "%s", res)
	} else {
		fmt.Fprintf(w, "L'étudiant avec l'id %d n'éxiste pas.", student.Id)
	}

}

//lgu.univ@gmail.com

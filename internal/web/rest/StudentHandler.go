package rest

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"io/ioutil"
	"net/http"

	"strconv"

	provider "internal/persistence"
	"internal/persistence/interfaces"

	"github.com/gorilla/mux"
)

var daoMemory interfaces.StudentDao = provider.GetDaoStudent()

//var daoMongodb ps.StudentDaoMongoDB = ps.NewStudentDaoMongo()

//var daoBolt ps.StudentDaoBolt = daoB.NewStudentDaoBolt()

// swagger:operation GET /students/{id} student studentsId
// ---
// summary: Return an Student provided by the id.
// description: If the Student is found, Student will be returned else Error Not Found (404) will be returned.
// parameters:
// - name: id
//   in: path
//   description: id of the language
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/studentRes"
//   "404":
//     "$ref": "#/responses/notFoundReq"

func StudentById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	student, err := daoMemory.Find(id)

	if err == nil {
		res, _ := json.Marshal(*student)
		fmt.Fprintf(w, "%s", res)
		return
	}

	fmt.Fprintf(w, "L'étudiant avec l'id %d n'éxiste pas.", id)
}

// swagger:operation GET /students/ student studentAll
// ---
// summary: Return students.
// description: Return all students.
// parameters:
// - None: None
// responses:
//   "200":
//     "$ref": "#/responses/studentRes"

func AllStudents(w http.ResponseWriter, r *http.Request) {

	res, _ := json.Marshal(daoMemory.FindAll())

	//res, _ := json.Marshal(daoMongodb.FindAll())

	//res, _ := json.Marshal(dao.FindAll())

	fmt.Fprintf(w, "%s", res)
}

// swagger:operation POST /students/ student studentCreate
// ---
// summary: Create a new student.
// description: If student creation is success, student will be returned with Created (201).
// parameters:
// - name: student
//   description: student to add to the list of students
//   in: body
//   required: true
//   schema:
//     "$ref": "#/internal/entities/Student"
// responses:
//   "200":
//     "$ref": "#/responses/okResp"

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var student entities.Student

	json.Unmarshal(reqBody, &student)

	if daoMemory.Create(student) {
		res, _ := json.Marshal(student)

		fmt.Fprintf(w, "%s", res)

		return
	}

	fmt.Fprintf(w, "L'étudiant existe déjà")

}

// swagger:operation DELETE /students/{id} student deleteStudent
// ---
// summary: Delete requested student by student id.
// description: Depending on the students id, HTTP Status Not Found (404) or HTTP Status OK (200) may be returned.
// parameters:
// - name: id
//   in: path
//   description: students id
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/okResp"
//   "404":
//     "$ref": "#/responses/notFoundReq"

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	student, err := daoMemory.Find(id)

	if err != nil {
		fmt.Fprintf(w, "L'étudiant avec l'id %d, n'éxiste pas.", id)
		return
	}

	if daoMemory.Delete(id) {
		res, _ := json.Marshal(student)
		fmt.Fprintf(w, "%s", res)
	}
}

// swagger:operation PUT /students/ student studentUpdate
// ---
// summary: Update a new student.
// description: If student update is success, student will be returned with Created (201) else if code doesn't exist return Not Found (404).
// parameters:
// - name: student
//   description: student to update in the list of students
//   in: body
//   required: true
//   schema:
//     "$ref": "#/internal/entities/Student"
// responses:
//   "200":
//     "$ref": "#/responses/okResp"
//   "404":
//     "$ref": "#/responses/notFound"

func PutStudent(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var student entities.Student

	json.Unmarshal(reqBody, &student)

	if daoMemory.Update(student) {
		res, _ := json.Marshal(student)
		fmt.Fprintf(w, "%s", res)
	} else {
		fmt.Fprintf(w, "L'étudiant avec l'id %d n'éxiste pas.", student.Id)
	}

}

//lgu.univ@gmail.com

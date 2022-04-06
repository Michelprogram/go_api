package main

import (
	"fmt"
	"net/http"

	"internal/web/rest"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("🚀 Lancement de l'api sur le port 8080...")

	r := mux.NewRouter()

	// Language Handlers

	r.HandleFunc("/apiV1/languages/{code}", rest.LanguageByCode).Methods("GET")
	r.HandleFunc("/apiV1/languages", rest.AllLanguages).Methods("GET")

	r.HandleFunc("/apiV1/languages", rest.CreateLanguage).Methods("POST")
	r.HandleFunc("/apiV1/languages/{code}", rest.DeleteLanguage).Methods("DELETE")
	r.HandleFunc("/apiV1/languages", rest.PutLanguage).Methods("PUT")

	// Student Handlers

	r.HandleFunc("/apiV1/students/{id}", rest.StudentById).Methods("GET")
	r.HandleFunc("/apiV1/students", rest.AllStudents).Methods("GET")

	r.HandleFunc("/apiV1/students", rest.CreateStudent).Methods("POST")
	r.HandleFunc("/apiV1/students/{id}", rest.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/apiV1/students", rest.PutStudent).Methods("PUT")

	http.ListenAndServe(":8080", r)

	fmt.Println("Fermeture de l'api ...")
}

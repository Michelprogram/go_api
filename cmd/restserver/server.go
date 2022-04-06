package main

import (
	"fmt"
	"net/http"

	"internal/entities"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("🚀 Lancement de l'api sur le port 8080...")

	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/language", testLanguage)

	http.ListenAndServe(":8080", r)

	fmt.Println("Fermeture de l'api ...")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Bonjour depuis l'api")
}

func testLanguage(w http.ResponseWriter, r *http.Request) {

	var language entities.Language = entities.NewLanguage("21", "Golang")

	fmt.Printf("Coucou from testLanguage : %s", language.String())

	fmt.Fprintf(w, "Test : %s", language.String())
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("🚀 Lancement de l'api ...")

	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8080", r)

	fmt.Println("Fermeture de l'api ...")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bonjour depuis l'api")
}

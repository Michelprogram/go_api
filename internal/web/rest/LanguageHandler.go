package rest

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"net/http"

	"github.com/gorilla/mux"
)

var languages []entities.Language = []entities.Language{
	entities.NewLanguage("21", "Go"), entities.NewLanguage("12", "Python"),
}

func LanguageByCode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	for _, element := range languages {
		if element.Code == vars["code"] {
			res, _ := json.Marshal(element)

			fmt.Fprintf(w, "%s", res)
			return
		}
	}

	fmt.Fprintf(w, "Le code %s n'a pas été trouvé.", vars["code"])
}

func AllLanguages(w http.ResponseWriter, r *http.Request) {

	res, _ := json.Marshal(languages)

	fmt.Fprintf(w, "%s", res)
}

func CreateLanguage(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Posts : ")

	fmt.Fprintf(w, "Done")
}

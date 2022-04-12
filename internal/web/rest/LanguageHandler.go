package rest

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	ps "internal/persistence"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var daoL ps.LanguageDao = ps.NewLanguageDaoMemory()

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

	reqBody, _ := ioutil.ReadAll(r.Body)
	var language entities.Language

	json.Unmarshal(reqBody, &language)

	languages = append(languages, language)

	res, _ := json.Marshal(language)

	fmt.Fprintf(w, "%s", res)
}

func PutLanguage(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var language entities.Language

	json.Unmarshal(reqBody, &language)

	for index, element := range languages {
		if element.Code == language.Code {
			languages[index] = language

			fmt.Fprintf(w, "%s", language)
			return
		}
	}

	fmt.Fprintf(w, "Le code %s n'est pas enregistré.", language.Code)

}

func DeleteLanguage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	for index, element := range languages {
		if element.Code == vars["code"] {

			res, _ := json.Marshal(element)

			languages = append(languages[:index], languages[index+1:]...)

			fmt.Fprintf(w, "%s", res)
			return
		}
	}

	fmt.Fprintf(w, "Le code %s n'a pas été trouvé.", vars["code"])
}

//lgu.univ@gmail.com

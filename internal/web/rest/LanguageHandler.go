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

//var daoL ps.LanguageDao = ps.NewLanguageDaoBolt()

func LanguageByCode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]

	language, err := daoL.Find(code)

	if err != nil {
		fmt.Fprintf(w, "Le code %s n'a pas été trouvé.", code)
		return
	}

	res, _ := json.Marshal(language)
	fmt.Fprintf(w, "%s", res)
}

func AllLanguages(w http.ResponseWriter, r *http.Request) {

	res, _ := json.Marshal(daoL.FindAll())

	fmt.Fprintf(w, "%s", res)
}

func CreateLanguage(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var language entities.Language

	json.Unmarshal(reqBody, &language)

	if daoL.Create(language) {
		res, _ := json.Marshal(language)

		fmt.Fprintf(w, "%s", res)
	} else {
		fmt.Fprintf(w, "Le Language existe déjà")
	}
}

func PutLanguage(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var language entities.Language

	json.Unmarshal(reqBody, &language)

	if daoL.Update(language) {
		res, _ := json.Marshal(language)
		fmt.Fprintf(w, "%s", res)
	} else {
		fmt.Fprintf(w, "Le code %s n'est pas enregistré.", language.Code)
	}
}

func DeleteLanguage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	code := vars["code"]

	if daoL.Exists(code) {
		if daoL.Delete(code) {
			fmt.Fprintf(w, "Le code %s a été supprimé.", code)
		}
	} else {
		fmt.Fprintf(w, "Le code %s n'a pas été trouvé.", code)
	}

}

//lgu.univ@gmail.com

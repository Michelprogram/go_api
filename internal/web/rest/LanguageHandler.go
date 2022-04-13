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

// swagger:operation GET /languages/{code} language languageCode
// ---
// summary: Return an Language provided by the code.
// description: If the Language is found, language will be returned else Error Not Found (404) will be returned.
// parameters:
// - name: code
//   in: path
//   description: code of the language
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/languageRes"
//   "404":
//     "$ref": "#/responses/notFoundReq"

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

// swagger:operation GET /languages/ language languageAll
// ---
// summary: Return languages.
// description: Return all languages.
// parameters:
// - None: None
// responses:
//   "200":
//     "$ref": "#/responses/languageRes"

func AllLanguages(w http.ResponseWriter, r *http.Request) {

	res, _ := json.Marshal(daoL.FindAll())

	fmt.Fprintf(w, "%s", res)
}

// swagger:operation POST /languages/ language languageCreate
// ---
// summary: Create a new language.
// description: If language creation is success, language will be returned with Created (201).
// parameters:
// - name: language
//   description: language to add to the list of languages
//   in: body
//   required: true
//   schema:
//     "$ref": "#/internal/entities/Language"
// responses:
//   "200":
//     "$ref": "#/responses/okResp"
//   "400":
//     "$ref": "#/responses/badReq"

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

// swagger:operation PUT /languages/ language languageUpdate
// ---
// summary: Update a new language.
// description: If language update is success, language will be returned with Created (201) else if code doesn't exist return Not Found (404).
// parameters:
// - name: language
//   description: language to update in the list of languages
//   in: body
//   required: true
//   schema:
//     "$ref": "#/internal/entities/Language"
// responses:
//   "200":
//     "$ref": "#/responses/okResp"
//   "404":
//     "$ref": "#/responses/notFound"

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

// swagger:operation DELETE /languages/{code} language deleteLanguage
// ---
// summary: Delete requested language by language code.
// description: Depending on the language code, HTTP Status Not Found (404) or HTTP Status OK (200) may be returned.
// parameters:
// - name: code
//   in: path
//   description: language code
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/okResp"
//   "404":
//     "$ref": "#/responses/notFoundReq"

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

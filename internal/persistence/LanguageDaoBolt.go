package persistence

import (
	"encoding/json"
	"errors"
	"internal/entities"

	"internal/persistence/bolt"
)

type LanguageDaoBolt struct {
}

var _ LanguageDao = (*LanguageDaoBolt)(nil)

var boltLanguage bolt.MyBolt = bolt.NewMyBolt()

func NewLanguageDaoBolt() LanguageDaoBolt {

	boltLanguage.CreateDatabase()

	return LanguageDaoBolt{}
}

func (d LanguageDaoBolt) Find(code string) (*entities.Language, error) {

	var language entities.Language

	res := boltLanguage.Get("Languages", code)

	if res == "" {
		return nil, errors.New("Le code n'existe pas")
	}

	json.Unmarshal([]byte(res), &language)

	return &language, nil
}

func (d LanguageDaoBolt) Exists(code string) bool {

	res := boltLanguage.Get("Languages", code)

	if res == "" {
		return false
	}

	return true

	//TODO : Revoir GET PB avec nombre qui commence par même chiffre 2 et insertion 21

}

func (d LanguageDaoBolt) Delete(code string) bool {

	err := boltLanguage.Delete("Languages", code)

	if err != nil {
		return false
	}

	return true
}

func (d LanguageDaoBolt) FindAll() []entities.Language {

	var languages []entities.Language

	for _, language := range boltLanguage.GetAll("Languages") {
		var lg entities.Language
		json.Unmarshal([]byte(language), &lg)
		languages = append(languages, lg)
	}

	return languages
}

func (d LanguageDaoBolt) Create(language entities.Language) bool {

	res, _ := json.Marshal(language)

	if !d.Exists(language.Code) {
		boltLanguage.Put("Languages", language.Code, string(res))
		return true
	}

	return false

}

func (d LanguageDaoBolt) Update(language entities.Language) bool {

	if d.Exists(language.Code) {

		res, _ := json.Marshal(language)

		boltLanguage.Delete("Languages", language.Code)

		boltLanguage.Put("Languages", language.Code, string(res))

		return true
	}

	return false
}

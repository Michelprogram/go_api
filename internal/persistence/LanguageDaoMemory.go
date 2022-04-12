package persistence

import (
	"errors"
	"internal/entities"
)

var languages []entities.Language = []entities.Language{
	entities.NewLanguage(2, "FR", "France"),
	entities.NewLanguage(1, "DE", "Allemagne"),
	entities.NewLanguage(3, "CH", "Chine"),
}

type LanguageDaoMemory struct {
}

var _ LanguageDao = (*LanguageDaoMemory)(nil)

func NewLanguageDaoMemory() LanguageDaoMemory {
	return LanguageDaoMemory{}
}

func (l LanguageDaoMemory) FindAll() []entities.Language {

	var newLanguages []entities.Language = languages

	for i := 0; i < len(newLanguages)-1; i++ {
		j := i + 1

		stA := newLanguages[i]
		stB := newLanguages[j]

		if stA.Id > stB.Id {

			newLanguages[i], newLanguages[j] = stB, stA
		}
	}

	return newLanguages
}

func (l LanguageDaoMemory) Find(code string) (*entities.Language, error) {

	for _, language := range languages {
		if language.Code == code {
			return &language, nil
		}
	}

	return nil, errors.New("Le code n'existe pas")
}

func (l LanguageDaoMemory) Exists(code string) bool {

	for _, language := range languages {
		if language.Code == code {
			return true
		}
	}

	return false
}

func (l LanguageDaoMemory) Delete(code string) bool {

	for index, language := range languages {
		if language.Code == code {
			languages = append(languages[:index], languages[index+1:]...)
			return true
		}
	}

	return false

}

func (l LanguageDaoMemory) Update(language entities.Language) bool {

	if !l.Exists(language.Code) {

		for index, element := range languages {
			if language.Id == element.Id {
				languages[index] = language
				return true
			}
		}
	}

	return false
}

func (l LanguageDaoMemory) Create(language entities.Language) bool {

	if !l.Exists(language.Code) {
		languages = append(languages, language)
		return true
	}

	return false

}

package daomemory

import (
	"errors"
	"internal/entities"
	"internal/persistence/interfaces"
)

var languages []entities.Language = []entities.Language{
	entities.NewLanguage(2, "FR", "France memory"),
	entities.NewLanguage(1, "DE", "Allemagne memory"),
	entities.NewLanguage(3, "CH", "Chine memory"),
}

type LanguageDaoMemory struct {
}

var _ interfaces.LanguageDao = (*LanguageDaoMemory)(nil)

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

package entities

import "fmt"

type Language struct {
	Code string
	Name string
}

func NewLanguage(code string, name string) Language {
	return Language{code, name}
}

func (l Language) String() string {
	return fmt.Sprintf("Language %s, code %s", l.Name, l.Code)
}

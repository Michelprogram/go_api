package entities

import "fmt"

type Language struct {
	Id   int
	Code string
	Name string
}

func NewLanguage(Id int, code string, name string) Language {
	return Language{Id, code, name}
}

func (l Language) String() string {
	return fmt.Sprintf("Language %s, code %s, id %d", l.Name, l.Code, l.Id)
}

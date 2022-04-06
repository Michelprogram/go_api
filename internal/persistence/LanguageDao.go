package persistence

import "internal/entities"

type LanguageDao interface {
	FindAll() []entities.Language
	Find(code int) *entities.Language
	Exists(code int) bool
	Delete(code int) bool
	Create(language entities.Language) bool
	Update(language entities.Language) bool
}

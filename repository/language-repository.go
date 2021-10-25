package repository

type LanguagePort interface {
	GetLanguage(index *int) (interface{}, error)
	GetLanguages() (interface{}, error)
	InsertLanguage(data interface{}) (interface{}, error)
	UpdateLanguage(index *int, data interface{}) (interface{}, error)
	DeleteLanguage(index *int) (interface{}, error)
}


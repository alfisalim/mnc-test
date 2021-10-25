package language

import "mnc-test/model"

type LanguagePort interface {
	GetLanguage(index *int) *model.CustomResponseHTTP
	GetLanguages() *model.CustomResponseHTTP
	InsertLanguage(req *model.Biodata) *model.CustomResponseHTTP
	UpdateLanguage(index *int, req *model.Biodata) *model.CustomResponseHTTP
	DeleteLanguage(index *int) *model.CustomResponseHTTP
}

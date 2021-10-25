package adapter

import (
	"mnc-test/model"
)

/**
Adapter layer used for build response from request, actually this layer
used for handling all about client, either third party, database, or etc.
*/

type LanguageAdapter struct{}

func NewLanguageAdapter() *LanguageAdapter {
	return &LanguageAdapter{}
}

func (l LanguageAdapter) GetLanguage(index *int) (interface{}, error) {
	result := model.Language[*index]
	return result, nil
}

func (l LanguageAdapter) GetLanguages() (interface{}, error) {
	result := model.Language
	return result, nil
}

func (l LanguageAdapter) InsertLanguage(data interface{}) (interface{}, error) {
	bind := data.(*model.Biodata)
	model.Language = append(model.Language, bind)
	return model.Language, nil
}

func (l LanguageAdapter) UpdateLanguage(index *int, data interface{}) (interface{}, error) {
	bind := data.(*model.Biodata)
	model.Language[*index] = bind
	return model.Language[*index], nil
}

func (l LanguageAdapter) DeleteLanguage(index *int) (interface{}, error) {
	model.Language = append(model.Language[:*index], model.Language[*index+1:]...)
	return model.Language, nil
}

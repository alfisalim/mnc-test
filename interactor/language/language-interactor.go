package language

import (
	"log"
	"mnc-test/model"
	"mnc-test/repository"
	"mnc-test/util"
	"net/http"
)

/**
	Interactor layer handling all about logic on this service,
	and then forward ther clearly request to adapter layers.
	this layers also used as receiver return from adapter layers
*/

type LanguageInteractor struct {
	LanguageRepo repository.LanguagePort
}

func NewLanguageInteractor(languageRepo repository.LanguagePort) *LanguageInteractor{
	return &LanguageInteractor{
		LanguageRepo: languageRepo,
	}
}

func (l LanguageInteractor) GetLanguage(index *int) *model.CustomResponseHTTP {
	if index == nil {
		log.Println("request is nil")
		response := util.ResponseBuilder(http.StatusBadRequest, "request is nil", nil)
		return response
	}

	// handling of when request greather than existing data
	count := len(model.Language)
	if *index >= count {
		log.Println("data not found")
		response := util.ResponseBuilder(http.StatusNotFound, "data not found", nil)
		return response
	}

	res, err := l.LanguageRepo.GetLanguage(index)
	if err != nil {
		log.Println("some error from adapter")
		response := util.ResponseBuilder(http.StatusInternalServerError, "some error form adapter", nil)
		return response
	}

	response := util.ResponseBuilder(http.StatusOK, "Successfully get specific data", res)
	return response
}

func (l LanguageInteractor) GetLanguages() *model.CustomResponseHTTP  {
	res, err := l.LanguageRepo.GetLanguages()
	if err != nil {
		log.Println("some error from adapter")
		response := util.ResponseBuilder(http.StatusInternalServerError, "some error form adapter", nil)
		return response
	}

	response := util.ResponseBuilder(http.StatusOK, "Successfully get all data", res)
	return response
}

func (l LanguageInteractor) InsertLanguage(req *model.Biodata) *model.CustomResponseHTTP {
	if req == nil {
		log.Println("request is nil")
		response := util.ResponseBuilder(http.StatusBadRequest, "request is nil", nil)
		return response
	}

	res, err := l.LanguageRepo.InsertLanguage(req)
	if err != nil {
		log.Println("some error from adapter")
		response := util.ResponseBuilder(http.StatusInternalServerError, "some error form adapter", nil)
		return response
	}

	response := util.ResponseBuilder(http.StatusCreated, "Successfully insert data", res)
	return response
}

func (l LanguageInteractor) UpdateLanguage(index *int, req *model.Biodata) *model.CustomResponseHTTP {
	if index == nil || req == nil {
		log.Println("request is nil")
		response := util.ResponseBuilder(http.StatusBadRequest, "request is nil", nil)
		return response
	}

	res, err := l.LanguageRepo.UpdateLanguage(index, req)
	if err != nil {
		log.Println("some error from adapter")
		response := util.ResponseBuilder(http.StatusInternalServerError, "some error form adapter", nil)
		return response
	}

	response := util.ResponseBuilder(http.StatusOK, "Successfully update data", res)
	return response
}

func (l LanguageInteractor) DeleteLanguage(index *int) *model.CustomResponseHTTP {
	if index == nil {
		log.Println("request is nil")
		response := util.ResponseBuilder(http.StatusBadRequest, "request is nil", nil)
		return response
	}

	res, err := l.LanguageRepo.DeleteLanguage(index)
	if err != nil {
		log.Println("some error from adapter")
		response := util.ResponseBuilder(http.StatusInternalServerError, "some error form adapter", nil)
		return response
	}

	response := util.ResponseBuilder(http.StatusOK, "Successfully delete data", res)
	return response
}
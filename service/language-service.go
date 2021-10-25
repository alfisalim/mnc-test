package service

import (
	"github.com/labstack/echo"
	"log"
	"mnc-test/interactor/language"
	"mnc-test/model"
	"mnc-test/util"
	"net/http"
	"strconv"
)

/**
	Service layer used for catch each request client
	and then forward the request to interactor layers.
	this layers also used as receiver return from interactor layers
*/

type LanguageService struct {
	Echo 		 *echo.Echo
	LanguagePort language.LanguagePort
}

func NewLanguageService(languagePort language.LanguagePort, echo *echo.Echo) *LanguageService {
	return &LanguageService{
		LanguagePort: languagePort,
		Echo: echo,
	}
}

func (l LanguageService) GetLanguage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res := l.LanguagePort.GetLanguage(&id)
	return c.JSON(res.StatusCode, res)
}

func (l LanguageService) GetLanguages(c echo.Context) error {
	res := l.LanguagePort.GetLanguages()
	return c.JSON(res.StatusCode, res)
}

func (l LanguageService) InsertLanguage(c echo.Context) error {
	bodyRes := new(model.Biodata)
	if err := c.Bind(bodyRes); err != nil {
		response := util.ResponseBuilder(http.StatusInternalServerError, err.Error(), nil)
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	res := l.LanguagePort.InsertLanguage(bodyRes)
	return c.JSON(res.StatusCode, res)
}

func (l LanguageService) UpdateLanguage(c echo.Context) error {
	bodyRes := new(model.Biodata)
	if err := c.Bind(bodyRes); err != nil {
		response := util.ResponseBuilder(http.StatusInternalServerError, err.Error(), nil)
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	id, _ := strconv.Atoi(c.Param("id"))

	res := l.LanguagePort.UpdateLanguage(&id, bodyRes)
	return c.JSON(res.StatusCode, res)
}

func (l LanguageService) DeleteLanguage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res := l.LanguagePort.DeleteLanguage(&id)
	return c.JSON(res.StatusCode, res)
}
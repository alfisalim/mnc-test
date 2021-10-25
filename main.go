package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"mnc-test/adapter"
	"mnc-test/interactor/language"
	"mnc-test/interactor/palindrome"
	"mnc-test/model"
	"mnc-test/rest"
	"mnc-test/rest/middlewareCustom"
	"mnc-test/service"
	"mnc-test/util"
	"net/http"
)

var Language []*model.Biodata

func main() {
	e := echo.New()

	e.Use(middleware.Logger()) //add middleware logging each request http
	e.Use(middleware.Recover()) // add middleware recover for recover panic anywhere in chain

	//base endpoint group for handle only GET http method
	eg := e.Group("/base")
	eg.Use(echo.WrapMiddleware(middlewareCustom.GETMethodValidation)) //call the function for validate method

	eg.GET("", func(context echo.Context) error {
		response := util.ResponseBuilder(http.StatusOK, "Hello Go developers", nil)
		return context.JSON(http.StatusOK, response)
	})

	//palindrome
	palindromeAdapter := adapter.NewPalindromeAdapter()
	palindromeInteractor := palindrome.NewPalindromeInteractor(palindromeAdapter)
	palindromeService := service.NewPalindromeService(palindromeInteractor, eg) //use echo group for get middleware validatoin

	//language
	languageAdapter := adapter.NewLanguageAdapter()
	languageInteractor := language.NewLanguageInteractor(languageAdapter)
	languageService := service.NewLanguageService(languageInteractor, e)


	server := rest.NewInitRouter(*languageService, *palindromeService)
	server.Handler()


	//eg.GET("palindrome", func(context echo.Context) error {
	//	text := context.QueryParam("text")
	//
	//	response := util.ResponseBuilder(http.StatusOK, "Palindrome", text)
	//
	//	ok := util.CheckPalindrom(text)
	//	if ok {
	//		return context.JSON(http.StatusOK, response)
	//	}
	//
	//	response.Message = "Not Palindrome"
	//	return context.JSON(http.StatusBadRequest, response)
	//})

	e.Logger.Fatal(e.Start(":7001"))
}

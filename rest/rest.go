package rest

import (
	"mnc-test/service"
)

type InitRouter struct {
	Language   service.LanguageService
	Palindrome service.PalindromeService
}

func NewInitRouter(language service.LanguageService, palindrome service.PalindromeService) *InitRouter {
	return &InitRouter{
		Language:   language,
		Palindrome: palindrome,
	}
}

func (i InitRouter) Handler() {
	//language router
	i.Language.Echo.GET("/language/:id", i.Language.GetLanguage)
	i.Language.Echo.GET("/languages", i.Language.GetLanguages)
	i.Language.Echo.POST("/language", i.Language.InsertLanguage)
	i.Language.Echo.PUT("/language/:id", i.Language.UpdateLanguage)
	i.Language.Echo.DELETE("/language/:id", i.Language.DeleteLanguage)

	//palindrome router
	i.Palindrome.EchoGroup.GET("/palindrome", i.Palindrome.PalindromeValidation)
}

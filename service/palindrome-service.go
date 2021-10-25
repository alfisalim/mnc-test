package service

import (
	"github.com/labstack/echo"
	"mnc-test/interactor/palindrome"
)

type PalindromeService struct {
	EchoGroup  		*echo.Group
	PalindromePort 	palindrome.PalindromePort
}

func NewPalindromeService(palindromePort palindrome.PalindromePort, echoGroup *echo.Group) *PalindromeService {
	return &PalindromeService{
		PalindromePort: palindromePort,
		EchoGroup: echoGroup,
	}
}

func (p PalindromeService) PalindromeValidation(c echo.Context) error  {
	text := c.QueryParam("text")
	res := p.PalindromePort.PalindromeValidation(text)
	return c.JSON(res.StatusCode, res)
}
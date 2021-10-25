package palindrome

import (
	"log"
	"mnc-test/model"
	"mnc-test/repository"
	"mnc-test/util"
	"net/http"
)

type PalindromeInteractor struct {
	PalindromeRepo repository.PalindromeRepository
}

func NewPalindromeInteractor(palindromeRepo repository.PalindromeRepository) *PalindromeInteractor {
	return &PalindromeInteractor{
		PalindromeRepo: palindromeRepo,
	}
}

func (p PalindromeInteractor) PalindromeValidation(req string) *model.CustomResponseHTTP {
	if req == "" {
		log.Println("request is nil")
		response := util.ResponseBuilder(http.StatusBadRequest, "request is nil", nil)
		return response
	}

	res := p.PalindromeRepo.PalindromeValidation(req)
	if res {
		response := util.ResponseBuilder(http.StatusOK, "Palindrome", req)
		return response
	}

	response := util.ResponseBuilder(http.StatusBadRequest, "Not Palindrome", req)
	return response
}
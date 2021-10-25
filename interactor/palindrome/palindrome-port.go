package palindrome

import "mnc-test/model"

type PalindromePort interface {
	PalindromeValidation(req string) *model.CustomResponseHTTP
}

package adapter

import (
	"mnc-test/util"
)

/**
Adapter layer used for build response from request, actually this layer
used for handling all about client, either third party, database, or etc.
this layers also used as receiver return from adapter layers
*/

type PalindromeAdapter struct{}

func NewPalindromeAdapter() *PalindromeAdapter {
	return &PalindromeAdapter{}
}

func (p PalindromeAdapter) PalindromeValidation(text string) bool {
	ok := util.CheckPalindrom(text)
	if ok {
		return true
	}

	return false
}

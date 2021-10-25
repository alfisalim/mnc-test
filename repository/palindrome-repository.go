package repository

type PalindromeRepository interface {
	PalindromeValidation(text string) bool
}

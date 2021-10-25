package util

import (
	"log"
)

func CheckPalindrom(word string) bool {
	countWord := len(word)
	opposite := countWord-1

	if countWord % 2 == 1 {
		for i:=0; i<=int(countWord/2)-1; i++ {
			if string(word[i]) != string(word[opposite]) {
				log.Println("Palindrom not found")
				return false
			}
			opposite--
		}

		log.Println("***Palindrom Founded***")
		return true
	}

	log.Println("Palindrom not found")
	return false
}

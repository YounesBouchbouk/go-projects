package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {

	clearedStr := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	left := 0
	right := len(clearedStr) - 1

	for left < right {

		if clearedStr[left] != clearedStr[right] {
			return false
		}

		left++
		right--
	}

	return true
}

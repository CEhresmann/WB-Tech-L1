package main

import (
	"fmt"
	"strings"
)

func hasUniqueCharacters(s string) bool {
	s = strings.ToLower(s)

	charSet := make(map[rune]struct{})

	for _, char := range s {
		if _, exists := charSet[char]; exists {
			return false
		}
		charSet[char] = struct{}{}
	}

	return true
}

func main() {
	testStrings := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
		"unique",
		"WBTech",
		"123456",
		"123aBcC",
	}

	for _, str := range testStrings {
		result := hasUniqueCharacters(str)
		fmt.Printf("Строка: \"%s\" — %t\n", str, result)
	}
}

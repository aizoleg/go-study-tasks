package main

import (
	"fmt"
	"unicode"
)

// AreCharsUnique проверяет, все ли символы в строке уникальные
// Функция является регистронезависимой
func AreCharsUnique(s string) bool {
	seen := make(map[rune]bool)

	for _, char := range s {
		lowerChar := unicode.ToLower(char) // Преобразуем символ в нижний регистр
		if seen[lowerChar] {
			return false // символ уже присутствует
		}
		seen[lowerChar] = true
	}
	return true
}

func main() {
	examples := []string{
		"abcd",
		"abCdefAaf",
		" aabcd",
	}

	for _, example := range examples {
		fmt.Printf("%s: %v\n", example, AreCharsUnique(example))
	}
}

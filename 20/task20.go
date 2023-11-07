package main

import (
	"fmt"
	"strings"
)

// ReverseWords переворачивает слова в строке
func ReverseWords(s string) string {
	words := strings.Split(s, " ") // Разбиваем строку на слова

	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ") // Объединяем слова обратно в строку
}

func main() {
	input := "snow dog sun"
	reversed := ReverseWords(input)
	fmt.Printf("Оригинал: %s\n Перевернутое: %s\n", input, reversed)
}

package main

import (
	"fmt"
)

// Reverse переворачивает строку
func Reverse(s string) string {
	runes := []rune(s) // Преобразуем строку в срез рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Меняем местами руны
	}
	return string(runes) // Преобразуем обратно в строку и возвращаем
}

func main() {
	input := "главрыба"
	reversed := Reverse(input)
	fmt.Printf("Оригинал: %s\n Перевернутое: %s\n", input, reversed)
}

package main

import (
	"fmt"
	"strings"
)

var justString string

// createHugeString создает большую строку из повторяющихся символов "A"
func createHugeString(size int) string {
	return strings.Repeat("A", size)
}

func someFunc() {
	v := createHugeString(1 << 10) // создаем строку длиной 1024 символа
	justString = string(v[:100])   // сохраняем только первые 100 символов
}

func main() {
	someFunc()
	fmt.Println("Длина justString:", len(justString))
}

package main

import "fmt"

func setBit(n int64, position uint, value bool) int64 {
	if value {
		// Устанавливаем i-й бит в 1 (устанавливаем бит)
		return n | (1 << position)
	} else {
		// Устанавливаем i-й бит в 0 (сбрасываем бит)
		return n &^ (1 << position)
	}
}

func main() {
	var num int64 = 42  // Пример исходного числа
	position := uint(2) // Номер бита, который мы хотим установить (считается с 0)
	value := true       // Установить бит в 1 (true) или 0 (false)

	result := setBit(num, position, value)

	fmt.Printf("Исходное число: %d\n", num)
	fmt.Printf("Установка бита %d в %v: %d\n", position, value, result)
}

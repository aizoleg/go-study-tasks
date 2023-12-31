package main

import (
	"fmt"
)

// swap меняет местами значения a и b без использования временной переменной
// Это достигается за счет арифметических операций сложения и вычитания
func swap(a, b int) (int, int) {
	a = a + b // a становится суммой a и b.
	b = a - b // (a + b) - b дает исходное значение a, так что b теперь равно a
	a = a - b // (a + b) - a дает исходное значение b, так что a теперь равно b

	return a, b
}

func main() {
	x := 5
	y := 10

	// Выводим исходные значения переменных
	fmt.Printf("До замены: x = %d, y = %d\n", x, y)

	// Меняем местами значения переменных x и y
	x, y = swap(x, y)

	// Выводим значения переменных после замены
	fmt.Printf("После замены: x = %d, y = %d\n", x, y)
}

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создание переменных типа big.Int с значениями > 2^20
	a := new(big.Int).SetInt64(1 << 21)       // 2^21
	b := new(big.Int).SetInt64((1 << 20) + 1) // 2^20 + 1

	// Результаты операций
	sum := new(big.Int)
	diff := new(big.Int)
	product := new(big.Int)
	quotient := new(big.Int)

	// Выполнение арифметических операций
	sum.Add(a, b)
	diff.Sub(a, b)
	product.Mul(a, b)
	quotient.Div(a, b)

	// Вывод результатов
	fmt.Printf("a = %s\n", a)
	fmt.Printf("b = %s\n", b)
	fmt.Printf("a + b = %s\n", sum)
	fmt.Printf("a - b = %s\n", diff)
	fmt.Printf("a * b = %s\n", product)
	fmt.Printf("a / b = %s\n", quotient)
}

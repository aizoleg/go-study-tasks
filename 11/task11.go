package main

import (
	"fmt"
)

// Intersect возвращает пересечение двух слайсов
func Intersect(a, b []int) []int {
	// Создаем карту для первого слайса
	m := make(map[int]bool)
	for _, item := range a {
		m[item] = true
	}

	// Проверяем элементы второго слайса
	var result []int
	for _, item := range b {
		if m[item] { // Если элемент присутствует и в первом слайсе, добавляем его в результат
			result = append(result, item)
			delete(m, item) // Удаляем элемент из карты, чтобы избежать дубликатов в результате
		}
	}

	return result
}

func main() {
	a := []int{1, 4, 2, 3, 5}
	b := []int{5, 6, 4, 7, 8}

	res := Intersect(a, b)
	fmt.Println("Пересечение множеств:", res)
}

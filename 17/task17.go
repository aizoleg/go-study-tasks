package main

import (
	"fmt"
	"sort"
)

func main() {

	// Неотсортированный срез
	numbers := []int{42, 12, 8, 15, 23, 16}

	// Сортировка среза
	sort.Ints(numbers)

	// Значение для поиска
	x := 15

	// Выполнение бинарного поиска
	index := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] >= x
	})

	// Проверка, найдено ли значение
	if index < len(numbers) && numbers[index] == x {
		fmt.Printf("Число %d найдено на позиции %d\n", x, index)
	} else {
		fmt.Printf("Число %d не найдено\n", x)
	}
}

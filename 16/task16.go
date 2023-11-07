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

	// Вывод отсортированного среза
	fmt.Println(numbers)
}

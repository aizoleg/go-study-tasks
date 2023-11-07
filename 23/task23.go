package main

import "fmt"

// RemoveAtIndex удаляет элемент по индексу из слайса
func RemoveAtIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		// Недопустимый индекс
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	s := []int{1, 2, 3, 4, 5}

	// Допустим, нам нужно удалить элемент с индексом 2
	updatedSlice := RemoveAtIndex(s, 2)

	fmt.Println(updatedSlice) // Вывод: [1 2 4 5]
}

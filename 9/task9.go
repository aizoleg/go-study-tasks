package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем каналы для отправки и получения данных
	inputChannel := make(chan int)
	outputChannel := make(chan int)
	var wg sync.WaitGroup

	// Горутина для записи чисел в первый канал
	go func() {
		defer close(inputChannel)
		numbers := []int{1, 2, 3, 4, 5}
		for _, num := range numbers {
			inputChannel <- num
		}
	}()

	// Горутина для умножения чисел на 2 и отправки во второй канал
	go func() {
		defer close(outputChannel)
		for num := range inputChannel {
			result := num * 2
			outputChannel <- result
		}
	}()

	// Горутина для вывода данных из второго канала в stdout
	go func() {
		defer wg.Done()
		for result := range outputChannel {
			fmt.Printf("Результат: %d\n", result)
		}
	}()

	wg.Add(1)

	// Ожидаем завершения всех горутин
	wg.Wait()
}

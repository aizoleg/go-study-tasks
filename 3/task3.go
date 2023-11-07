package main

import (
	"fmt"
	"sync"
)

// Функция для вычисления суммы квадратов чисел
func calculateSquareSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num * num
	}
	return sum
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	numGoroutines := 2 // Количество горутин для параллельных вычислений
	resultChan := make(chan int, numGoroutines)

	chunkSize := len(numbers) / numGoroutines
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == numGoroutines-1 {
			end = len(numbers)
		}

		// Запускаем горутину для обработки части чисел
		go func(start, end int) {
			defer wg.Done()
			partialSum := calculateSquareSum(numbers[start:end])
			fmt.Printf("Часть суммы: %d\n", partialSum) // Выводим часть суммы для каждой горутины
			resultChan <- partialSum                    // Отправляем результат в канал
		}(start, end)
	}

	go func() {
		wg.Wait()
		close(resultChan) // Закрываем канал после завершения всех горутин
	}()

	totalSum := 0
	for partialSum := range resultChan {
		totalSum += partialSum
	}

	fmt.Printf("Сумма квадратов чисел: %d\n", totalSum)
}

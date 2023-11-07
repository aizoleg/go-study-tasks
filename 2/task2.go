package main

import (
	"fmt"
	"sync"
)

func calculateSquare(wg *sync.WaitGroup, num int) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения функции

	square := num * num
	fmt.Printf("%d^2 = %d\n", num, square)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup // Создаем WaitGroup для ожидания завершения горутин

	for _, num := range numbers {
		wg.Add(1) // Увеличиваем счетчик WaitGroup перед запуском горутины
		go calculateSquare(&wg, num)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}

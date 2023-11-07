package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем map для хранения данных
	data := make(map[string]int)
	// Создаем мьютекс для синхронизации доступа к map
	var mutex sync.Mutex

	var wg sync.WaitGroup
	numWrites := 100

	for i := 0; i < numWrites; i++ {
		wg.Add(1)
		go func(key string, value int) {
			defer wg.Done()
			// Захватываем мьютекс для безопасной записи в map
			mutex.Lock()
			data[key] = value
			// Освобождаем мьютекс
			mutex.Unlock()
		}(fmt.Sprintf("key%d", i), i)
	}

	wg.Wait()

	// Вывод данных из map
	mutex.Lock()
	for key, value := range data {
		fmt.Printf("%s: %d\n", key, value)
	}
	// Освобождаем мьютекс
	mutex.Unlock()
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter структура счетчика
type Counter struct {
	value int64
}

// Increment инкрементирует счетчик на 1
func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Value возвращает текущее значение счетчика
func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	const numGoroutines = 1000

	// Запускаем 1000 горутин, каждая из которых инкрементирует счетчик 1000 раз
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}

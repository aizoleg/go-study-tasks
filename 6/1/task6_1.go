package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Горутина завершена.")
				return
			default:
				fmt.Println("Горутина работает.")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	done <- true // Отправляем сигнал завершения
	time.Sleep(1 * time.Second)
}

/* Использование канала для сигнала завершения горутины
Прием сигнала завершения в горутине может привести к выходу из нее */

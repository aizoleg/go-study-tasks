package main

import (
	"fmt"
	"time"
)

func main() {
	stop := false

	go func() {
		for !stop {
			fmt.Println("Горутина работает.")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Горутина завершена.")
	}()

	time.Sleep(3 * time.Second)
	stop = true // Останавливаем выполнение горутины
	time.Sleep(1 * time.Second)
}

/*Использование переменной для управления выполнением:
также можно использовать общую переменную для управления выполнением горутины */

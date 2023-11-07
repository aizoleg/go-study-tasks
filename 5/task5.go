package main

import (
	"fmt"
	"time"
)

func main() {
	N := 5 // Количество секунд работы программы
	dataChannel := make(chan int)

	// Запускаем горутину для отправки данных в канал
	go func() {
		defer close(dataChannel)
		for i := 1; i <= N; i++ {
			dataChannel <- i
			time.Sleep(1 * time.Second) // Задержка в 1 секунду
		}
	}()

	// Чтение данных из канала
	for data := range dataChannel {
		fmt.Printf("Прочитано: %d\n", data)
	}

	fmt.Println("Программа завершена.")
}

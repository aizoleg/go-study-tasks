package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	dataChannel := make(chan int)
	numWorkers := 5 // Количество воркеров

	// Создаем канал для сигнала завершения программы
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	// Функция для воркера
	worker := func(id int) {
		for data := range dataChannel {
			fmt.Printf("Воркер %d: Прочитано из канала: %d\n", id, data)
		}
	}

	// Запускаем воркеры
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}

	// Главный поток записывает данные в канал
	go func() {
		for i := 1; ; i++ {
			dataChannel <- i
		}
	}()

	// Ожидание сигнала завершения
	<-signalChannel

	// Завершение работы всех воркеров
	close(dataChannel)
	wg.Wait()

	fmt.Println("Программа завершена.")
}

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена.")
				return
			default:
				fmt.Println("Горутина работает.")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel() // Отменяем выполнение горутины
	time.Sleep(1 * time.Second)
}

/*Использование контекста (context):
Стандартный пакет context в Go позволяет отменять выполнение горутин через контекст*/

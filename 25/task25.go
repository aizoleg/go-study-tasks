package main

import (
	"fmt"
	"time"
)

// MySleep "засыпает" на указанную длительность
func MySleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Начало работы...")

	// "Засыпаем" на 2 секунды
	MySleep(2 * time.Second)

	fmt.Println("Пробуждение после сна!")
}

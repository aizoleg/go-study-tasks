package main

import (
	"fmt"
)

// determineType определяет и печатает тип значения, сохраненного в интерфейсе
func determineType(name string, variable interface{}) {
	switch v := variable.(type) {
	case int:
		fmt.Printf("Переменная '%s' имеет тип: int\n", name)
	case string:
		fmt.Printf("Переменная '%s' имеет тип: string\n", name)
	case bool:
		fmt.Printf("Переменная '%s' имеет тип: bool\n", name)
	case chan int: // канал типа int
		fmt.Printf("Переменная '%s' имеет тип: channel\n", name)
	default:
		fmt.Printf("Переменная '%s' имеет неизвестный тип: %T\n", name, v)
	}
}

func main() {
	var myInt int
	var myString string
	var myBool bool
	var myChan chan int

	determineType("myInt", myInt)
	determineType("myString", myString)
	determineType("myBool", myBool)
	determineType("myChan", myChan)
}

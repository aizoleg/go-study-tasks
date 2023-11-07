package main

import "fmt"

// Определение структуры Human
type Human struct {
	Name string
	Age  int
}

// Метод, принадлежащий структуре Human
func (h *Human) Speak() {
	fmt.Printf("%s говорит: Привет, мой возраст %d лет.\n", h.Name, h.Age)
}

// Определение структуры Action с встраиванием структуры Human
type Action struct {
	Human      // Встраивание структуры Human в структуру Action
	ActionName string
}

// Метод, принадлежащий структуре Action
func (a *Action) DoAction() {
	fmt.Printf("%s выполняет действие: %s\n", a.Name, a.ActionName)
}

func main() {
	// Создаем объект Human
	person := Human{
		Name: "Иван",
		Age:  30,
	}

	// Создаем объект Action
	action := Action{
		Human:      person, // Встраивание объекта Human
		ActionName: "Прыгать",
	}

	// Вызываем методы Speak() и DoAction()
	action.Speak()    // Метод Speak() унаследован от Human
	action.DoAction() // Метод DoAction() принадлежит структуре Action
}

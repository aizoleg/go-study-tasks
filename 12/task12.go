package main

import (
	"fmt"
)

// StringsSet реализует множество строк
type StringsSet map[string]bool

// NewStringsSet создает множество из последовательности строк
func NewStringsSet(items ...string) StringsSet {
	set := make(StringsSet)
	for _, item := range items {
		set[item] = true
	}
	return set
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := NewStringsSet(sequence...)

	fmt.Println("Множество строк:", set)
}

// использование map устранит дубликаты, так как в map каждый ключ может встречаться только один раз

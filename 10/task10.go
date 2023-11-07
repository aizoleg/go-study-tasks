package main

import (
	"fmt"
	"math"
)

func main() {
	// Исходная последовательность температур
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем карту для группировки температурных значений
	temperatureGroups := make(map[int][]float64)

	// Группируем температуры
	for _, temp := range temperatures {
		group := int(math.Ceil(temp/10)) * 10 //используем math.Ceil для округления до ближайшего целого вверх
		temperatureGroups[group] = append(temperatureGroups[group], temp)
	}

	// Выводим группы температур
	for group, temps := range temperatureGroups {
		fmt.Printf("%d: %v\n", group, temps)
	}
}

package main

import (
	"fmt"
	"math"
)

// Point представляет точку на плоскости с координатами x и y
type Point struct {
	x float64
	y float64
}

// NewPoint - конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// DistanceTo вычисляет расстояние от текущей точки до другой точки p
func (p *Point) DistanceTo(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	// Создание двух точек с использованием конструктора
	p1 := NewPoint(1, 1)
	p2 := NewPoint(4, 5)

	// Вычисление и вывод расстояния между двумя точками
	fmt.Println(p1.DistanceTo(p2))
}

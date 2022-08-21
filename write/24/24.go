package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64 // с маленькой буквы - потому что инкапсулированы
}

// конструктор:
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// вычисляем расстояние по формуле:
func Distance(a, b Point) float64 {
	return math.Sqrt((b.x-a.x)*(b.x-a.x) + (b.y-a.y)*(b.y-a.y))
}

func main() {
	// задаем точки:
	a := NewPoint(1, 1)
	b := NewPoint(2, 2)
	fmt.Println(Distance(a, b)) // получаем результат
}

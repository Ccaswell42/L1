package main

import (
	"fmt"
)

// fmt работает через рефлексию. Рефлексия - работа рантайме. reflect.TypeOf() - возвращает тип данных.

//Функция принимает пустой интерфейс и через форматированный вывод (%T) пишет в консоль ее тип.
func printType(val interface{}) {
	fmt.Printf("Тип данных: %T\n", val)
}
func main() {
	var a int
	var b string
	var c bool
	d := make(chan struct{})
	printType(a)
	printType(b)
	printType(c)
	printType(d)
}

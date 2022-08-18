package main

import "fmt"

//В GO нет наслеодвания как такого, но есть структуры, в поле которое можно записать другую структуру,
//причем методы дочерних структур тоже наследуются. Если методы называются одинаково, реализация родительского метода
//будет переписана реализацие дочернего.

type Human struct {
	Name string
	Age  int
	Sex  string
}

// Метод Print структуры Human выводит в консоль поле Name
func (h Human) Print() {
	fmt.Println(h.Name)
}

// Метод PrintAge структуры Human выводит в консоль поле Age
func (h Human) PrintAge() {
	fmt.Println(h.Age)
}

type Action struct {
	Job     string
	Hobbies string
	Human
}

// Метод Print структуры Action (одинаковое название метода как и у стуктуры HUMAN) выводит в консоль поле Job
func (a Action) Print() {
	fmt.Println(a.Job)
}

// Метод PrintHobbies структуры Action выводит в консоль поле Hobbies
func (a Action) PrintHobbies() {
	fmt.Println(a.Hobbies)
}

func main() {
	var h = Human{"Dmitriy", 19, "male"}
	var a = Action{"Lawyer", "football", h}

	// тип HUMAN реализует свой метод Print и выводит имя
	h.Print()
	// тип ACTION реализует свой метод Print и выводит работу (используется дочерний, а не родительский метод,
	// потому что названия одинаковы)
	a.Print()
	// Из типа ACTION обращаемся к полю HUMAN, а затем к методу Print, принадлежащему типу HUMAN
	a.Human.Print()
	// Из типа ACTION обращаемся к полю HUMAN, а затем к методу PrintAge, принадлежащему типу HUMAN
	a.Human.PrintAge()
	// напрямую из типа ACTION обращаемся к методу PrintAge, который принадлежит типу HUMAN
	a.PrintAge()
}

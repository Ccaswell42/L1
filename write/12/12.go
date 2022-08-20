package main

import "fmt"

// множество - набор уникальных элементов

// функция для проверки наличия данного элемента в массиве.
// возвращает true если элемент есть в массиве, false -если нет.
func checkVal(res []string, val string) bool {
	for _, d := range res {
		if d == val {
			return true
		}
	}
	return false
}

// создаем массив, в который будем записывать элементы нужного нам множества.
// с помощью функции checkVal проверяем, добавляли мы уже такой элмент или нет. если нет - записываем в массив.
func makeArr(items []string) []string {
	var res []string

	for _, val := range items {
		if !(checkVal(res, val)) {
			res = append(res, val)
		}
	}
	return res
}
func main() {
	// данная нам последовательность строк
	items := []string{"cat", "cat", "dog", "cat", "tree"}
	// выводим полученное множество в консоль
	fmt.Println(makeArr(items))

}

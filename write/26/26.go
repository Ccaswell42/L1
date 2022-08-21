package main

import (
	"fmt"
	"os"
	"strings"
)

//  проверяем есть ли искомый символ в строке до оперделенной позиции: false- есть, true - нет
func CheckVal(arr []rune, val rune, i int) bool {
	for d := 0; d < i; d++ {
		if arr[d] == val {
			return false
		}
	}
	return true
}

func Checker(str string) bool {
	// чтобы проверка была регистронезависимой приведем всю строку в нижний регистр и сразу скастуем в слайс рун
	data := []rune(strings.ToLower(str))

	// в цикле проверим был ли этот символ в строке раньше
	for i, val := range data {
		if !CheckVal(data, val, i) { // если был - возвращаем false
			return false
		}
	}
	return true
}

func main() {
	// строку берем из аргумента коммандной строки
	fmt.Println(Checker(os.Args[1])) // выводим результат в консоль
}

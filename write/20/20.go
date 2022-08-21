package main

import (
	"fmt"
	"os"
	"strings"
)

func revers(str string) string {

	arr := strings.Split(str, " ") // с помощью функцции Split разбиваем строку по пробелу в массив

	var newArr []string

	// в цикле записываем наш массив в новый в обратном порядке
	for i := len(arr) - 1; i >= 0; i-- {
		newArr = append(newArr, arr[i])
	}
	// соедениям наш массив в одну строку через разделитьель "пробел" и возвращаем
	return strings.Join(newArr, " ")

}
func main() {
	// строку берем из аргумента коммандной строки
	fmt.Println(revers(os.Args[1])) // выводим результат в консоль

}

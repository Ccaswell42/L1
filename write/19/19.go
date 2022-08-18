package main

import (
	"fmt"
	"os"
)

func revers(new []rune) string {

	//в цикле разворачиваем срез рун на месте
	last := len(new) - 1
	for first := 0; first < len(new)/2; first++ {
		new[first], new[last] = new[last], new[first]
		last--

	}
	// кастуем срез рун к типу string и возвращаем
	return string(new)

}
func main() {
	// принимаем строку из аргумента коммандной строки и кастуем к типу срез рун
	fmt.Println(revers([]rune(os.Args[1])))
}

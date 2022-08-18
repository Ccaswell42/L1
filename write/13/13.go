package main

import "fmt"

// функция  revers принимает в качестве аргумента 2 числа, и возвращает их в обратном порядке

func revers(i, d int) (int, int) {
	return d, i
}
func main() {

	i := 0
	d := 1

	fmt.Println("i=", i, "d=", d)
	// самый простой способ
	i, d = d, i
	fmt.Println("i=", i, "d=", d)
	// с помощью функции revers
	i, d = revers(i, d)
	fmt.Println("i=", i, "d=", d)
	// математический способ
	i = i + d
	d = i - d
	i = i - d

	fmt.Println("i=", i, "d=", d)
}

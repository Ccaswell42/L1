package main

import (
	"fmt"
	"os"
)

func trim(str []rune) []string {
	var new []rune
	var arr []string

	for _, val := range str {
		if val != 32 {
			new = append(new, val)
		} else {
			arr = append(arr, string(new))
			new = []rune("")
			continue
		}
	}
	arr = append(arr, string(new))
	return arr
}

func revers(str string) string {
	arr := trim([]rune(str))
	last := len(arr) - 1
	var res string

	for first := 0; first < len(arr)/2; first++ {
		arr[first], arr[last] = arr[last], arr[first]
		last--
	}
	for i, val := range arr {

		if i != 0 {
			res += " "
		}
		res += val
	}
	return res
}

func main() {
	//fmt.Println([]rune(" "))
	str := revers(os.Args[1])
	fmt.Println(str)

}

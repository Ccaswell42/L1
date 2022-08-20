package main

import "fmt"

func binSearch(arr []int, val int) int {
	first := 0
	last := len(arr) - 1

	for {
		mid := (first + last) / 2
		if val < arr[mid] {
			last = mid - 1
		} else if val > arr[mid] {
			first = mid + 1
		} else {
			return mid
		}
		if first > last {
			return -1
		}
	}
}
func main() {

	arr := []int{-8, 1, 19, 24, 29, 76, 99}
	fmt.Println(binSearch(arr, -8))
	fmt.Println(binSearch(arr, 24))
	fmt.Println(binSearch(arr, 76))
	fmt.Println(binSearch(arr, 90))

}

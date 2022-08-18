package main

import (
	"fmt"
)

//функция delete принимает в качесте варгументов срез и элемент,который нужно удалить
func delete(arr []int, i int) []int {
	// просто удалить i-й элемент из среза нельзя, можно лишь создать новый. с
	//сделаем это с помощью функции append
	arr = append(arr[0:i], arr[i+1:]...)
	return arr
}

func main() {
	arr := []int{12, 14, 15, 16, 17, 18, 19}
	fmt.Println(delete(arr, 3))

}

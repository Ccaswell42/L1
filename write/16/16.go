package main

import "fmt"

func main() {
	arr := []int{19, 27, 14, 0, 3, -9, 155, 1884, -21475, 0, 3, 9}
	pr := QuickSort(arr)

	fmt.Println(pr)

}

func QuickSort(arr []int) []int {

	//  Создаем подмассивы для элементов, меньше опорного; для элементов больше опорного; для повторяющихся элементов;
	// для вывода результата;

	var less []int
	var high []int
	var base []int
	var ret []int

	if len(arr) < 2 {
		return arr
	}
	// в качестве опорного элемента выберем серединный
	pivot := arr[len(arr)/2]

	//в цикле мтерируемся по заданному массиву, элементы которые меньше опорного элемента записываем в массив less,
	// а элементы, которые больше в массив high, опорный элемент - в массив base.
	for _, val := range arr {
		if val == pivot {
			base = append(base, pivot)
			continue
		}
		if val < pivot {
			less = append(less, val)
		} else {
			high = append(high, val)
		}
	}
	//Сообираем результирующий массив, сначала идет подмассив less, затем опорный элемент, и в конце подмассив High,
	// так же рекурсивно применяем наш алгоритм к соответствующим подмассивам.
	ret = append(ret, QuickSort(less)...)
	ret = append(ret, base...)
	ret = append(ret, QuickSort(high)...)
	return ret
}

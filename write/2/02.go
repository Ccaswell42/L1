package main

import (
	"fmt"
	"sync"
)

// В цикле вызываем горутины, которые принимают числа и печатают квадраты этих чисел в консоль,
//используем  Waitgroup, для того чтобы в основной горутине дождаться выполнения всех остальных,
// а не выйти недождавшись
func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, i := range arr {
		go func(i int) {
			fmt.Println(i * i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

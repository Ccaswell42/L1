package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	arr := []int{2, 4, 6, 8, 10}
	useAtomic(arr)
	useMutex(arr)
	useChannels(arr)

}

/*
В цикле запускаем горутины, в которых конкуренто высчитываются квадраты элементов массива.
С помощью примитива синхронизации atomic атомарно инкрементируем число на результат данных вычислений,
в результате получим сумму квадратов чисел среза.
Используем waitGroup для того чтобы дождаться результат выполнения всех горутин

*/
func useAtomic(arr []int) {
	var res int64 = 0
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, i := range arr {
		go func(c int) {
			atomic.AddInt64(&res, int64(c*c))
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Println(res)
}

/*
В цикле запускаем горутины, в которых высчитываем квадрат числа и прибавляем к результатам вычислений других горутин
в переменную "res", для избежания гонки данных используем мьютексы (lock/unlock).
Используем waitGroup для того чтобы дождаться результат выполнения всех горутин.
*/

func useMutex(arr []int) {
	var res int = 0
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, i := range arr {
		go func(c int) {
			mu.Lock()
			res += c * c
			mu.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Println(res)
}

/*
В первом цикле запускаем горутины, в которых высчитываем квадрат числа и отправлем в канал,
во втором цикле читаем из канала и складыаем результат вместе.
*/
func useChannels(arr []int) {

	ch := make(chan int)

	for _, i := range arr {
		go func(c int, ch chan int) {
			ch <- c * c
		}(i, ch)
	}

	res := 0
	for i, _ := range arr {
		i = <-ch
		res += i
	}
	fmt.Println(res)

}

/* результаты бенчмарка:
useAtomic: 3751 наносекунд, 12 аллокаций памяти, 264 байта на операцию
UseMutex: 4044 наносекунд, 13 аллокаций памяти, 312 байт на операцию
useChannels: 5411 наносекунд, 6 аллокаций памяти, 256 байт на операцию
*/

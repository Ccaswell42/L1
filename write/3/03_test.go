package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

/*
	go test -bench . -benchmem 03_test.go
*/

func BenchmarkAtomic(b *testing.B) {
	arr := []int{2, 4, 6, 8, 10}
	for i := 0; i < b.N; i++ {
		if x := useAtomic(arr); x != 220 {
			b.Fatalf("Unexpected string: %d", x)
		}
	}
}

func BenchmarkMutex(b *testing.B) {
	arr := []int{2, 4, 6, 8, 10}
	for i := 0; i < b.N; i++ {
		if x := useMutex(arr); x != 220 {
			b.Fatalf("Unexpected string: %d", x)
		}
	}
}

func BenchmarkChannels(b *testing.B) {
	arr := []int{2, 4, 6, 8, 10}
	for i := 0; i < b.N; i++ {
		if x := useChannels(arr); x != 220 {
			b.Fatalf("Unexpected string: %d", x)
		}
	}
}

func useAtomic(arr []int) int64 {
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
	//fmt.Println(res)
	return res
}

func useMutex(arr []int) int {
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
	//fmt.Println(res)
	return res
}

func useChannels(arr []int) int {

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
	//fmt.Println(res)
	return res

}

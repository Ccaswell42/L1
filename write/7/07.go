package main

import (
	"fmt"
	"sync"
)

func writer(counters map[int]int, wg *sync.WaitGroup, c int, mu *sync.Mutex) {
	for i := 0; i < 5; i++ {
		mu.Lock()
		counters[c*10+i] = 1
		mu.Unlock()
	}
	wg.Done()
}

func main() {
	counters := make(map[int]int)
	wg := &sync.WaitGroup{}
	wg.Add(5)
	mu := &sync.Mutex{}
	for i := 0; i < 5; i++ {
		go writer(counters, wg, i, mu)
	}
	wg.Wait()
	fmt.Println("counters result", counters)
}

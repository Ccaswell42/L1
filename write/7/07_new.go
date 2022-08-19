package main

import (
	"fmt"
	"sync"
)

func writer(counters sync.Map, wg *sync.WaitGroup, c int) {
	for i := 0; i < 5; i++ {
		counters.Store(c*10+i, 1)
	}
	wg.Done()
}

func main() {
	var counters sync.Map
	wg := &sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go writer(counters, wg, i)
	}
	wg.Wait()
	fmt.Println("counters result", counters.)

}

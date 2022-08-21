package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Наша структура-счетчик, где лежит некое число и sync.Mutex для синхронизации конкурентных операций
type Counter struct {
	num int64
	sync.Mutex
}

// Метод инкрементирования счетчика с использованием mutex
func (c *Counter) IncMutex() {
	c.Lock()
	c.num++
	c.Unlock()

}

// Метод инкрементирования счетчика с использованием atomic
func (c *Counter) IncAtomic() {
	atomic.AddInt64(&c.num, 1)
}

func main() {
	c := 150                 // количество воркеров и насколько нужно инкрементировать число
	cntr := &Counter{num: 0} // инициализируем нашу структуру
	wg := &sync.WaitGroup{}  // waitGroup для того чтобы дождаться выполнения горутин
	wg.Add(c)
	// в цикле запускаем горутины, которые будут инкрементровать наш счетчик
	for i := 0; i < c; i++ {
		go func(counter *Counter, gw *sync.WaitGroup) {
			cntr.IncMutex() // можно выбрать другой метод с использованием Atomic
			gw.Done()
		}(cntr, wg)
	}
	wg.Wait() // дожидаемся выполнения наших горутин

	fmt.Println(cntr.num) // выводим итоговое значение счетчика

}

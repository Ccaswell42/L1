package main

import (
	"fmt"
	"sync"
)

//конкурентная запись в map 2-мя способами: с помощью mutex и с помощью типа sync.Map

/*
1) С помощью mutex. Функция writer в цикле записывает в map(переданную в качестве аргумента) некоторые данные,
обособляя запись специальными методами (lock/unlock). Так же используем waitGroup для того, чтобы дождаться конца
работы всех наших горутин.

*/
func writer(counters map[int]int, wg *sync.WaitGroup, c int, mu *sync.Mutex) {
	for i := 0; i < 5; i++ {
		mu.Lock()
		counters[c*10+i] = 1
		mu.Unlock()
	}
	wg.Done()
}

func useMutex() {
	counters := make(map[int]int) // создем map
	wg := &sync.WaitGroup{}       // создаем waitGroup
	wg.Add(5)                     //инкрементируем счетчик waitGroup
	mu := &sync.Mutex{}           // создаем Mutex
	// В цикле запускаем нескольго горутин, которы конкурентно будут записывать в нашу map некоторые данные,
	// так же передаем mutex и  waitGroup
	for i := 0; i < 5; i++ {
		go writer(counters, wg, i, mu)
	}
	wg.Wait()                            // дожидаемся завершения работы наших горутин
	fmt.Println("with mutex:", counters) // выводим нашу map  в консоль

}

/*
2) С помощью типа sync.Map. Похож на обычную map, но является потокобезопасным для одновременного использования
несколькими горутинами.
В фукнции writerSyncMap записываем в нашу sync.Map некоторые данные с помощбю специального метода Store.
Так же используем waitGroup для того, чтобы дождаться конца работы всех наших горутин.
*/
func writerSyncMap(counters *sync.Map, wg *sync.WaitGroup, c int) {
	for i := 0; i < 5; i++ {
		counters.Store(c*10+i, 1)
	}
	wg.Done()
}

func useSyncMap() {
	counters := &sync.Map{} // инициализируем sync.Map
	wg := &sync.WaitGroup{} // создаем waitGroup
	wg.Add(5)               //инкрементируем счетчик waitGroup
	// В цикле запускаем нескольго горутин, которы конкурентно будут записывать в нашу sync.Map некоторые данные,
	// так же передаем waitGroup
	for i := 0; i < 5; i++ {
		go writerSyncMap(counters, wg, i)
	}
	wg.Wait() // дожидаемся завершения работы наших горутин

	// выводим данные из нашей sync.Map в консоль с помощью специального метода Range
	fmt.Println("sync.Map:")
	counters.Range(func(k, v interface{}) bool {
		fmt.Printf("[%d:%d] ", k, v)
		return true
	})
	fmt.Println()
}

func main() {
	useMutex() // пример с использованием mutex
	fmt.Println()
	useSyncMap() // пример с использованием sync.Map

}

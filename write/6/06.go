package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// все возможные способы остановки горутин

// 1) остановка горутины с помощью контекста (например withCancel)
// в функции byContext реализуем бесконечный цикл, который постоянно выводит в консоль некоторую информацию.
// при вызове  cancelFunc в основной горутине - прекращает работу
func byContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("finish work")
			return
		default:
			fmt.Println("do some work")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// 2)Остановка горутины с помощью канала
// в функции byChannel реализуем бесконечный цикл, который постоянно выводит в консоль некоторую информацию.
// // при отправке данных в канал в основной горутине - прекращает работу.
// пустая структура ничего не весит
func byChannel(cancelCh chan struct{}) {
	for {
		select {
		case <-cancelCh:
			fmt.Println("finish work")
			return
		default:
			fmt.Println("do some work")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// 3) Остановка горутины с помощью waitGroup
// функция byWaitgroup принимает указатель на waitGroup, делает некоторую работу, затем ожидает дикрементации счетчика,
// после чего завершает работу

func byWaitGroup(wg *sync.WaitGroup) {

	for i := 0; i < 11; i++ {
		fmt.Println(i, ") do some work")
	}
	wg.Wait()
	fmt.Println("finish work")
}

// 4) Остановка горутины с помощью закрытия канала
// функция byCloseChannel принимает канал, и в цикле читает данные из канала. когда канал закрывается, цикл прекращается
// и горутина завершает свою работу
func byCloseChannel(ch chan string) {
	for val := range ch {
		fmt.Println(val)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("finish work")
}
func main() {
	/*
		1 способ
	*/
	//создаем контекст
	ctx, finish := context.WithCancel(context.Background())

	// запускаем горутины, которая останавливается при помощи CancelContext
	go byContext(ctx)
	// ждем некоторое время пока отработает
	time.Sleep(5 * time.Second)
	// вызываем функцию отмены, горутина завершается
	finish()
	//
	/*
		2 способ
	*/
	// создаем канал с пустой стурктурой, пустая структура ничего не весит
	cancelCh := make(chan struct{})
	// запускаем горутину, в которую передаем наш канал
	go byChannel(cancelCh)
	// ждем некоторое время пока отработает
	time.Sleep(5 * time.Second)
	// отправляем  в канал пустую структуру, горутина завершает работу
	cancelCh <- struct{}{}

	/*
		3 способ
	*/
	// создаем waitGroup по указателю
	wg := &sync.WaitGroup{}
	// инкрементируем счетчик на 1
	wg.Add(1)
	// запускаем нашу горутину
	go byWaitGroup(wg)
	// ждем некоторое время пока отработает
	time.Sleep(5 * time.Second)
	// декрмементируем наш счетчик, после чего наша горутина прекратит ждать и завершит работу
	wg.Done()

	/*
		4 способ
	*/
	// создаем канал
	ch := make(chan string)

	// запускаем нашу горутину и передаем ей канал
	go byCloseChannel(ch)
	// в цикле пишем в канал некоторую информацию
	for i := 0; i < 11; i++ {
		ch <- "do some work"
	}
	// закрываем канал, после чего наша горутина прекращает работу
	close(ch)
}

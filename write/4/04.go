package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

// функция worker принимает канал и читет из него и пишет в stdout пока канал открыт, далее
// декрменетирует счетчик waitgroup и завершает работу
func worker(i int, ch chan string, wg *sync.WaitGroup) {
	for data := range ch {
		fmt.Printf("worker_%d: %s\n", i, data)
	}
	fmt.Printf("worker_%d finished\n", i)
	wg.Done()
}

// функция loop принимает контекст и канал для отправки информации воркерам,
//в бесконечном цикле по умолчанию пишет информацию в канал воркерам, если срабатывает кансел контекст
// прекращает работу
func loop(ctx context.Context, ch chan string) {

	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- "some_information"
		}
	}
}

func main() {
	// количество воркеров берем из аргумента коммандной строки
	c, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// создаем waigroup чтобы дождаться работы всех воркеров и не выйти из главной горутины до этого
	wg := &sync.WaitGroup{}
	// создаем канал, в который будет бесконечно записываться информация, и откуда будут читать наши воркеры
	ch := make(chan string, 1)

	// создаем кансел контекст для того чтобы использоваьт его, когда придет сигнал "Ctrl+C" и
	//завершить работу наших воркеров
	cntx, finishFunc := context.WithCancel(context.Background())

	// в отдельной горутине создаем наших воркеров и инкрементируем счетчик waitgroup
	go func() {
		for i := 0; i < c; i++ {
			wg.Add(1)
			go worker(i, ch, wg)

		}
	}()
	// для обработки сигнала " Ctrl+C" будем использовать пакет signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)
	// в отдельной горутине будем ждать сигнала, и как только получим, запустим функциию отмены нашего контекста
	// и работа нашей программы завершится
	go func() {
		<-quit
		finishFunc()
	}()
	// запускам наш бесконечный чикл
	loop(cntx, ch)
	// дожидаемя завершения работы наших воркеров
	wg.Wait()
	fmt.Println("bye bye")
}

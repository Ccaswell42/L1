package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// фнукция wtiter в бесконечном цикле  пишет значения в канал, как только сработает таймер контекста -
// закроет канал и выйдет из функции
func writer(ctx context.Context, ch chan int) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- i
		}
		time.Sleep(100 * time.Millisecond) //  для наглядности
		i++
	}

}
func main() {

	// количество секунд, после которых, программа завершится возьмем из аргумента коммандной строки
	sec, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// устанавливаем контекст с таймаутом
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)

	// создаем канал куда будем писать и откуда читать
	ch := make(chan int, 1)
	// запускаем в отдельной горутине нашу запись
	go writer(ctx, ch)
	// читаем пока есть что читать, и пока канал не закрыт
	for val := range ch {
		fmt.Println(val)
	}
	// как только сработает таймаут, запись прекратится, чтение тоже, и мы заврешим работу нашей прогрмамы
	fmt.Println("Finish")

}

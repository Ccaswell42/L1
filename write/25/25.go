package main

import (
	"fmt"
	"time"
)

func fnSleep(t time.Duration) {
	<-time.After(t) // функция отправляет в канал текущее время после того, как истечет указанное в аргументе время

}
func main() {
	fmt.Println("start:", time.Now()) // сравним время начала и конца таймера
	fnSleep(10 * time.Second)         // установим наш sleep на 10 секунд
	fmt.Println("finish:", time.Now())

}

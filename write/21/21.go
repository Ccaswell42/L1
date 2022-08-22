package main

import (
	"errors"
	"fmt"
	"log"
)

// Идея паттерна "адаптер" заключается в том, что для того чтобы уже существующая структура и ее методы удовлетворяли
// интерфейсу не нужно переписывать структуру или методы, а нужно сделать для нее адаптер.

// Есть структура Wallet для которой мы сделаем адаптер
type Wallet struct {
	cash int
}

// Структура Card
type Card struct {
	balance int
}

// Метод Pay для структуры Card
func (c *Card) Pay(amount int) error {
	if amount > c.balance {
		return errors.New("не хватает денег на карте")
	}
	c.balance -= amount
	return nil
}

// есть интерфейс Payyer, которому удовлетворяет только структура Card
type Payer interface {
	Pay(int) error
}

// Для того чтобы интерфейсу Payer удовлетворяла и структура Wallet сделаем новую структуру WalletAdapter,
// которая унаследует все от структуры Wallet + будет иметь свой собственный метод Pay, который будет удовлетворять
// нашему интерфейсу Payer
type WalletAdapter struct {
	*Wallet
}

// Метод Pay структуры WallettAfapter удовлетворяет интерфейсу Payer
func (w *WalletAdapter) Pay(amount int) error {
	if amount > w.cash {
		return errors.New("не хватает денег в кошельке")
	}
	w.cash -= amount
	return nil
}

// Функция Buy принимает интерфейс Payer;  Списывает деньги с карты или с кошелька;
// Dыдает ошибку, если денег не достаточно
func Buy(p Payer) {
	if err := p.Pay(100); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Покупка совершена упешно с помощью %T\n", p)
}

func main() {
	//инициализируем наши стуктуры
	myWallet := &WalletAdapter{&Wallet{cash: 1000}}
	myCard := &Card{1000}

	// используем функцию Buy для обоих объектов
	Buy(myCard)
	Buy(myWallet)
}

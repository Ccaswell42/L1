package main

import "errors"

// есть 2 структуры: старая и новая. у старой есть какие то методы
// есть интерфейс который реализует новая структура. чтобы этот интерфейс реализовал старую струтуру
// нужно сделать адаптер который унаследует старую структуру и будет иметь метод который будет соответствовать данному
// интерфейсу.

type Wallet struct {
	cash int
}

type Card struct {
	balance int
}

func (c *Card) Pay(amount int) error {
	if amount > c.balance {
		return errors.New("не хватает денег на карте")
	}
	c.balance -= amount
	return nil
}

type Payer interface {
	Pay(int) error
}

type WalletAdapter struct {
	Wallet
}

func (w *WalletAdapter) Pay(amount int) error {
	if amount > w.cash {
		return errors.New("не хватает денег в кошельке")
	}
	w.cash -= amount
	return nil
}

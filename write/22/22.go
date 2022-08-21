package main

import (
	"fmt"
	"math/big"
)

func main() {

	// для операций над большими числами используем пакет Big
	c := big.NewInt(9223372036854775807) // можно использовать не больше чем INT64max
	d := big.NewInt(9223372036854775807)

	// числа, больше, чем INT64max можно задать через строки таким образом:
	a, o := new(big.Int).SetString("-1800000000000000000000000", 10)
	b, p := new(big.Int).SetString("2200000000000000000000000", 10)

	if !o || !p {
		fmt.Println("bigInt задан некорректно")
		return
	}

	MUL := new(big.Int).Mul(c, d) // умножение
	fmt.Println("сумма с и d:", MUL)
	DIV := new(big.Int).Div(c, d) // деление
	fmt.Println("деление с на d:", DIV)
	SUM := new(big.Int).Add(a, b) // сумма
	fmt.Println("сумма a и b:", SUM)
	DIFF := new(big.Int).Sub(a, b) // вычитание
	fmt.Println("вычитание a из b:", DIFF)
}

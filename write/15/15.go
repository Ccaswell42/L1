package main

import (
	"fmt"
	"strings"
)

// Первая проблема - это использование строк. Для того чтобы "срезать" символы а не байты, лучше использовать руны.
// Вторая - это то, что после среза куска строки, наш срез будет указывать на нашу HugeString строку как на базовы массив,
// и тем самым будет продолжать держать ее в памяти.
var justString string

// делаем большую строку длиной n из одинаковых элементов. взял символ кириллица для наглядности, так как он весит 2 байта.
func createHugeString(n int) string {
	return strings.Repeat("ы", n)
}

func someFunc() {
	v := createHugeString(1 << 10)
	runeV := []rune(v)               // кастуем к рунам
	sliceV := make([]rune, 100, 100) // создаем слайс с cap 100
	copy(sliceV, runeV[:100])        // копируем а не срезаем, чтобы больше не ссылаться на огромный массив
	justString = string(sliceV)      // кастуем к типу строка
	fmt.Println(justString)
}

func main() {
	someFunc()
}

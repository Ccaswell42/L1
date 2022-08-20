package main

import "fmt"

/*
|: поразрядная дизъюнкция (операция ИЛИ или поразрядное сложение).
Возвращает 1, если хотя бы один из соответствующих разрядов обоих чисел равен 1
*/
/*
^: поразрядное исключающее ИЛИ. Возвращает 1, если только один из соответствующих разрядов обоих чисел равен 1
*/

// Чтобы получить нужный бит - сдвинем единицу на порядковый номер нужного нам бита влево (1 << count)
func changeBit(num, count int64, bit bool) int64 {
	switch bit {
	// Если нужно поменять бит на 1, то прменем поразрядное сложение
	case true:
		return num | (1 << count)
	// Если нужно поменять бит на ноль, то применем поразрадное исключающее ИЛИ
	default:
		return num ^ (1 << count)
	}

}
func main() {
	var num int64
	var count int64
	var bit bool

	fmt.Scanln(&num)                        // записываем из консоли число, в котором нужно поменять бит
	fmt.Scanln(&count)                      // записываем из консоли порядковый номер бита
	fmt.Scanln(&bit)                        // записываем из консоли бит (0 или 1)
	fmt.Println(changeBit(num, count, bit)) // выводим в консоль рещультат наших вычислений

}

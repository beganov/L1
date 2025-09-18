package main

import "fmt"

var a int = 3
var b int = 5

// L1.13
// Поменять местами два числа без использования временной переменной.
// Подсказка: примените сложение/вычитание или XOR-обмен.

func main() {
	fmt.Println(a, b)

	a, b = b, a
	fmt.Println(a, b) //средствами языка

	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b) //сложением

	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b) //XOR
}

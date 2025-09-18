package main

import (
	"fmt"
	"math/big"
)

// L1.21
// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).
// Комментарий: в Go тип int справится с такими числами,
// но обратите внимание на возможное переполнение для ещё больших значений.
//  Для очень больших чисел можно использовать math/big.

func main() {
	var a, b, c big.Int // используeтся math/big.
	a.SetInt64(1 << 21) // значения больше 1 миллиона
	b.SetInt64(1 << 20)
	fmt.Println("a: ", a.String())
	fmt.Println("b: ", b.String())
	fmt.Println("a+b: ", c.Add(&a, &b).String()) // сложение
	fmt.Println("a-b: ", c.Sub(&a, &b).String()) // вычитание
	fmt.Println("a*b: ", c.Mul(&a, &b).String()) // умножение
	fmt.Println("a/b: ", c.Div(&a, &b).String()) // деление
}

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b, c big.Int
	a.SetInt64(1 << 21)
	b.SetInt64(1 << 20)
	fmt.Println("a: ", a.String())
	fmt.Println("b: ", b.String())
	fmt.Println("a+b: ", c.Add(&a, &b).String())
	fmt.Println("a-b: ", c.Sub(&a, &b).String())
	fmt.Println("a*b: ", c.Mul(&a, &b).String())
	fmt.Println("a/b: ", c.Div(&a, &b).String())
}

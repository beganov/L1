package main

import "fmt"

var n int64 = 52          // изменяемая переменная
var i int64 = 2           // изменяемый бит
var bitValue bool = false // значение бита

func main() {
	if bitValue {
		fmt.Println(n | (1 << (i - 1)))
	} else {
		fmt.Println(n &^ (1 << (i - 1)))
	}
}

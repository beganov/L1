package main

import (
	"fmt"
	"strings"
)

// L1.14
// Разработать программу, которая в runtime способна определить тип переменной,
// переданной в неё (на вход подаётся interface{}).
// Типы, которые нужно распознавать: int, string, bool, chan (канал).
// Подсказка: оператор типа switch v.(type) поможет в решении.

type TestStruct struct { // для демонстрации работы с кастомными типами
	intTest    int
	stringTest string
	boolTest   bool
}

func main() {
	variables := []interface{}{
		52,
		"hello",
		true,
		make(chan int),
		make(chan string),
		make(chan bool),
		make(chan TestStruct),
		&TestStruct{},
		3.14,
		nil,
	}

	for _, i := range variables {
		fmt.Println(determineType(i))
	}

}
func determineType(value interface{}) string {
	switch value.(type) { // оператор типа switch v.(type)
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int, chan string, chan bool, chan interface{}:
		//не покрывает все каналы, так как каждый канал - разный тип
		return "chan"
	default:
		if strings.Split(fmt.Sprintf("%T", value), " ")[0] == "chan" {
			// поэтому имеет место рассмотреть каналы через "%T"
			return "chan"
		}
		return "unknown type"
	}
}

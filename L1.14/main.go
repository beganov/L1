package main

import (
	"fmt"
	"strings"
)

type TestStruct struct {
	intTest    int
	stringTest string
	boolTest   bool
}

func main() {
	variables := []interface{}{
		42,
		"hello",
		true,
		make(chan int),
		make(chan string),
		make(chan bool),
		make(chan TestStruct),
		3.14,
		nil,
	}

	for _, i := range variables {
		fmt.Println(determineType(i))
	}

}
func determineType(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int, chan string, chan bool, chan interface{}: //не покрывает все канал, каждый канал - разный тип
		return "chan"
	default:
		if strings.Split(fmt.Sprintf("%T", value), " ")[0] == "chan" {
			return "chan"
		}
		return "unknown type"
	}
}

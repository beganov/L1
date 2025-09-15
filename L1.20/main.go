package main

import (
	"fmt"
	"strings"
)

var mainString = "sun dog snow глав🐟рыба зам🐟 🐟рыба🐟"

func main() {
	// итерирование по рунам с флагом
	// мягко говоря не самый читаемый/идеоматичный способ
	// но в рамках данной задачи имеет место быть
	result := ""
	counter := 0
	flag := false
	for i, j := range mainString {
		if flag { //т.к. i - позиция в байтах,
			//разные руны разный размер имеют,
			// то пропускаем пробел буквально, не арифметически
			flag = false
			result = mainString[counter:i] + result
			counter = i
		}
		if j == ' ' {
			flag = true
		}
	}
	result = mainString[counter:] + " " + result
	fmt.Println(result)

	//Разбиение и склейка по пробелам пакетом strings
	splitString := strings.Split(mainString, " ")
	for i, j := 0, len(splitString)-1; i < len(splitString)/2; i, j = i+1, j-1 {
		(splitString)[i], (splitString)[j] = (splitString)[j], (splitString)[i]
	}
	fmt.Println(strings.Join(splitString, " "))
}

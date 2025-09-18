package main

import (
	"fmt"
	"strings"
)

// L1.26
// Разработать программу, которая проверяет, что все символы в строке встречаются один раз
//  (т.е. строка состоит из уникальных символов).
// Вывод: true, если все символы уникальны, false, если есть повторения.
//  Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.
// Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.
// Подумайте, какой структурой данных удобно воспользоваться для проверки условия.

var str = "sunS🐟"

func main() {
	variables := []string{
		"52",
		"hello",
		"sunS🐟",
		"Sun",
	}

	for _, i := range variables {
		fmt.Println(isUnique(i))
	}
}

func isUnique(str string) bool {
	reMap := make(map[rune]int)
	str = strings.ToLower(str)
	// Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.
	for _, j := range str {
		reMap[j]++
		if reMap[j] > 1 {
			return false
		}
	}
	return true
}

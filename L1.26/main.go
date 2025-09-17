package main

import (
	"fmt"
	"strings"
)

var str = "sunS🐟"

func main() {
	fmt.Println(isUnique(str))
}

func isUnique(str string) bool {
	reMap := make(map[rune]int)
	str = strings.ToLower(str)
	for _, j := range str {
		reMap[j]++
		if reMap[j] > 1 {
			return false
		}
	}
	return true
}

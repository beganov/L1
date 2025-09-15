package main

import "fmt"

var mainString = "главрыба🐟"

func main() {
	runeArr := []rune(mainString)
	for i := 0; i < len(runeArr)/2; i++ {
		runeArr[i], runeArr[len(runeArr)-1-i] = runeArr[len(runeArr)-1-i], runeArr[i]
	}
	mainString = string(runeArr)
	fmt.Println(mainString)
}

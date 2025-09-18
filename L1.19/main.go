package main

import "fmt"

// L1.19
// Разработать программу, которая переворачивает подаваемую на вход строку.
// Например: при вводе строки «главрыба» вывод должен быть «абырвалг».
// Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
// то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

//добавил эмодзи для демонстрации работы с другим размером байта
var mainString = "главрыба🐟"

func main() {
	runeArr := []rune(mainString)
	for i := 0; i < len(runeArr)/2; i++ { //проходим до середины строки
		runeArr[i], runeArr[len(runeArr)-1-i] = runeArr[len(runeArr)-1-i], runeArr[i]
		//меняем местами зеркальные относительно центра элементы
	}
	mainString = string(runeArr)
	fmt.Println(mainString)
}

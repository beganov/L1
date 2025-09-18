package main

import "fmt"

var array = []string{"cat", "cat", "dog", "cat", "tree"}

// L1.12
// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree").
// Создать для неё собственное множество.
// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func main() {
	var result []string
	mapA := make(map[string]int)

	// ищем уникальные значения и их количество
	for _, i := range array {
		mapA[i]++
	}

	// каждое найденное слово добавляем в выходное множество
	for j := range mapA {
		if mapA[j] != 0 {
			result = append(result, j)
		}
	}
	fmt.Println(result)
}

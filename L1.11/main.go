package main

import "fmt"

// L1.11
// Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) —
// т.е. вывести элементы, присутствующие и в первом, и во втором.
// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}

var arrayA = []int{3, 2, 1, 3}
var arrayB = []int{4, 3, 2, 3}

func main() {
	var result []int
	mapA := make(map[int]int)

	// считаем количество уникальных в первом массиве
	for _, i := range arrayA {
		mapA[i]++
	}

	// ищем пересечения с вторым массивом
	for _, j := range arrayB {
		if mapA[j] != 0 {
			result = append(result, j)
			mapA[j]-- // на случай повторяющихся значений
		}
	}
	fmt.Println(result)
}

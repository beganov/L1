package main

import "fmt"

// L1.17
// Реализовать алгоритм бинарного поиска встроенными методами языка.
// Функция должна принимать отсортированный слайс и искомый элемент,
// возвращать индекс элемента или -1, если элемент не найден.
// Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.

var array = []int{1, 3, 6, 7, 8, 55}

func main() {
	fmt.Println(binarySearch(array, 6))
}

func binarySearch(arr []int, searchedInt int) int {

	low := 0
	high := len(arr) - 1
	mid := high / 2
	for low <= high { // пока границы поиска не сошлись
		if arr[mid] == searchedInt {
			return mid // возвращаем индекс элемента
		} else {
			if searchedInt > arr[mid] { // если мы выше середины
				low = mid + 1 // двигаем нижнюю границу
			} else {
				high = mid - 1 // иначе верхнюю
			}
			mid = (low + high) / 2
		}
	}
	return -1 // -1, если элемент не найден.
}

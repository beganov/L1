package main

import "fmt"

// L1.15
// Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.
// Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
//  Для выбора опорного элемента можно взять середину или первый элемент.

var array = []int{10, 5, 3, 2, 4}

func main() {
	fmt.Println(quickSort(array))
}

func quickSort(arr []int) []int {
	if len(arr) < 2 { // базовый случай
		return arr
	} else {
		pivot := arr[0] // опорный - первый элемент
		var less = []int{}
		var greater = []int{}
		for _, num := range arr[1:] {
			if pivot > num { //распределяем массив на массивы больших и меньших опорного
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}

		// рекурсивно сортируем массивы
		less = append(quickSort(less), pivot)
		greater = quickSort(greater)
		return append(less, greater...)
	}
}

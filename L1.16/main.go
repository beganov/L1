package main

import "fmt"

var array = []int{10, 5, 3, 2, 4}

func main() {
	fmt.Println(quickSort(array))
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		var less = []int{}
		var greater = []int{}
		for _, num := range arr[1:] {
			if pivot > num {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}

		less = append(quickSort(less), pivot)
		greater = quickSort(greater)
		return append(less, greater...)
	}
}

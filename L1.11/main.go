package main

import "fmt"

var arrayA []int = []int{3, 2, 1, 3}
var arrayB []int = []int{4, 3, 2, 3}

func main() {
	var result []int
	mapA := make(map[int]int)
	for _, i := range arrayA {
		mapA[i]++
	}
	for _, j := range arrayB {
		if mapA[j] != 0 {
			result = append(result, j)
			mapA[j]-- // на случай повторяющихся значений
		}
	}
	fmt.Println(result)
}

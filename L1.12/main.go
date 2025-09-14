package main

import "fmt"

var array []string = []string{"cat", "cat", "dog", "cat", "tree"}

func main() {
	var result []string
	mapA := make(map[string]int)
	for _, i := range array {
		mapA[i]++
	}
	for j := range mapA {
		if mapA[j] != 0 {
			result = append(result, j)
		}
	}
	fmt.Println(result)
}

package main

import "fmt"

var array = []int{1, 3, 6, 7, 8, 55}

func main() {
	fmt.Println(binarySearch(array, 6))
}

func binarySearch(arr []int, searchedInt int) int {

	low := 0
	high := len(arr) - 1
	mid := high / 2
	for low <= high {
		if arr[mid] == searchedInt {
			return mid
		} else {
			if searchedInt > arr[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
			mid = (low + high) / 2
		}
	}
	return -1
}

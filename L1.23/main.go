package main

import "fmt"

var i = 3
var array = []int{1, 2, 3, 4, 5}

func main() {
	copy(array[i:], array[i+1:])
	array = array[:len(array)-1]
	fmt.Println(array)
}

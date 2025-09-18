package main

import (
	"fmt"
	"unsafe"
)

// L1.23
// Удалить i-ый элемент из слайса.
// Продемонстрируйте корректное удаление без утечки памяти.
// Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента
// (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.

var i = 3
var array = []int{1, 2, 3, 4, 5}

func main() {
	fmt.Println(array, unsafe.SliceData(array))
	copy(array[i:], array[i+1:])          // сдвинуть хвост слайса
	newSlice := make([]int, len(array)-1) //создаем новый слайс, чтобы отвязаться от старого
	copy(newSlice, array[:len(array)-1])
	array = newSlice //уменьшить длину слайса на 1.
	fmt.Println(array, unsafe.SliceData(array))
}

package main

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"
)

var justString string
var alsoJustString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] // //при выходе из функции
	// все еще ссылаемся на огромную строку v, что приведет к утечке памяти
	alsoJustString = strings.Clone(v[:100])
	//Из комментария к функции strings.Clone()
	// Clone returns a fresh copy of s.
	// It guarantees to make a copy of s into a new allocation,
	// which can be important when retaining only a small substring
	// of a much larger string. Using Clone can help such programs
	// use less memory. Of course, since using Clone makes a copy,
	// overuse of Clone can make programs use more memory.
	// Clone should typically be used only rarely, and only when
	// profiling indicates that it is needed.
	// For strings of length zero the string "" will be returned
	// and no allocation is made.

	justStringPtr := unsafe.StringData(justString)
	alsoJustStringPtr := unsafe.StringData(alsoJustString)
	vPtr := unsafe.StringData(v)

	fmt.Printf("original v string pointer: %p\n", vPtr)
	fmt.Printf("justString pointer: %p\n", justStringPtr)     //как видно, оригинал и слайс ссылаются на одну область
	fmt.Printf("alsoString pointer: %p\n", alsoJustStringPtr) //а clone нет
	runtime.GC()
}

func main() {
	someFunc()
}

func createHugeString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = 'a'
	}
	return string(result)
}

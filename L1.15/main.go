package main

import (
	"fmt"
	"strings"
	"unsafe"
)

// L1.15
// Рассмотреть следующий код и ответить на вопросы:
// к каким негативным последствиям он может привести и как это исправить?
// (утечка памяти, исправить можно использованием функции Clone (или до Go 1.18 string([]byte(v[:100]))))
// Приведите корректный пример реализации.
// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }
// func main() {
//   someFunc()
// }
// Вопрос: что происходит с переменной justString?
// (ссылается на ту же область памяти , что и исходная строка, GC не сможет очистить всю исходную строку)

var justString string
var alsoJustString string
var alsoString string

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
	alsoString = string([]byte(v[:100])) //постфактум нашел такой вариант

	justStringPtr := unsafe.StringData(justString)
	alsoJustStringPtr := unsafe.StringData(alsoJustString)
	alsoStringPtr := unsafe.StringData(alsoString)
	vPtr := unsafe.StringData(v)

	fmt.Printf("original v string pointer: %p\n", vPtr)
	fmt.Printf("justString pointer: %p\n", justStringPtr)         //как видно, оригинал и слайс ссылаются на одну область
	fmt.Printf("alsoJustString pointer: %p\n", alsoJustStringPtr) //а clone нет
	fmt.Printf("alsoString pointer: %p\n", alsoStringPtr)         //и string([]byte(v[:100])) тоже нет
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

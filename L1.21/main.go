package main

import "fmt"

// L1.21 \
// Реализовать паттерн проектирования «Адаптер» на любом примере.
// Описание: паттерн Adapter позволяет сконвертировать интерфейс
// одного класса в интерфейс другого, который ожидает клиент.
// Продемонстрируйте на простом примере в Go: у вас есть существующий интерфейс
// (или структура) и другой, несовместимый по интерфейсу потребитель
// — напишите адаптер, который реализует нужный интерфейс и делегирует вызовы к встроенному объекту.
// Поясните применимость паттерна, его плюсы и минусы, а также приведите реальные примеры использования.

type existSystem struct { // у вас есть существующий интерфейс (или структура)
}

type clientInterface interface { // и другой, несовместимый по интерфейсу потребитель
	doSomethingElse() string
}

func (es existSystem) doSomething() string {
	return "Do something ordinary"
}

func doSomethingElse() string {
	return "Do something else"
}

type Adapter struct { // адаптер, который реализует нужный интерфейс
	sys *existSystem
}

func (a Adapter) doSomethingElse() string {
	return a.sys.doSomething() // и делегирует вызовы к встроенному объекту.
}

func main() {
	adapter := Adapter{sys: &existSystem{}}
	var ci clientInterface = adapter
	fmt.Println(ci.doSomethingElse())
}

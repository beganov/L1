package main

import "fmt"

type existSystem struct {
}

type clientInterface interface {
	doSomethingElse() string
}

func (es existSystem) doSomething() string {
	return "Do something ordinary"
}

func doSomethingElse() string {
	return "Do something else"
}

type Adapter struct {
	sys *existSystem
}

func (a Adapter) doSomethingElse() string {
	return a.sys.doSomething()
}

func main() {
	adapter := Adapter{sys: &existSystem{}}
	var ci clientInterface = adapter
	fmt.Println(ci.doSomethingElse())
}

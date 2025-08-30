package main

import "fmt"

type Human struct {
	age  int
	name string
}

type Action struct {
	duration int
	Human
}

func (h *Human) GetAge() string {
	return fmt.Sprintf("Age: %d", h.age)
}

func (h *Human) GetName() string {
	return "My name is " + h.name
}

func (a *Action) GetDuration() string {
	return fmt.Sprintf("Action duration: %d", a.duration)
}

func main() {
	var act Action
	act.age = 52
	act.name = "Slime shady"
	act.duration = 1337

	fmt.Println(act.GetAge())  // метод  Human
	fmt.Println(act.GetName()) // метод Human
	fmt.Println(act.duration)  // метод Action
}

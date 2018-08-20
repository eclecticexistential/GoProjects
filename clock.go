package main

import "fmt"

type Minutes struct {
	value int // encapsulated so that other funcs cannot change value type
 }

func (m *Minutes) Increment(){
	m.value = (m.value + 1) % 60
}

func (m *Minutes) Set(newValue int) {
	m.value = newValue % 60
}

func (m *Minutes) Display(){
	fmt.Println(m.value)
}

func main() {
	minutes := Minutes()
	minutes.Set(63)
	for i := 1; i <= 3; i++{
		minutes.Increment()
		minutes.Display()
	}
}
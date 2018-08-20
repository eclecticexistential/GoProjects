package main

import "fmt"

type Minutes int

func (m *Minutes) Increment(){
	*m = (*m + 1) % 60
}

//* is pointer & is address

func main() {
	minutes := Minutes(58)
	for i := 1; i<=3; i++{
		minutes.Increment()
		fmt.Println(minutes)
	}
}
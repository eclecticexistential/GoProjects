package main

import (
	"fmt"
)

type MyType float64

func (m MyType) MyMethod(){
	fmt.Println("In function MyMethod on MyType declared")
}

func main(){
	myValue := MyType(1.23)
	myValue.MyMethod()
}
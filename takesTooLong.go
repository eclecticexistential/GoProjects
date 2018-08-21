package main

import (
	"fmt"
	"time"
)

func longTask() {
	fmt.Println("Starting long task")
	time.Sleep(3 * time.Second)
	fmt.Println("Long task finished")
}

func main() {
	longTask()
	longTask()
	longTask()
}
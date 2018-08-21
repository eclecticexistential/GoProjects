package main

import (
	"fmt"
	"time"
	"rand"
)

func longTask(channel chan int) {
	delay := rand.Intn(5)
	fmt.Println("Starting long task")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("Long task finished")
	channel <- delay
}

func main() {
	rand.Seed(time.Now().Unix())
	channel := make(chan int)
	for i := 1; i <= 3; i++ {
		go longTask(channel)
	}
	for i := 1; i <= 3; i++ {
		fmt.Println("Took", <-channel, "seconds")
	}
}
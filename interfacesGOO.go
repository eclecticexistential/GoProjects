package main

import "fmt"

type Monitor struct {
	Resolution 	string
	Connector 	string
	Value 		Float64
}

type HardDrive struct {
	Type string
	Connector string
	Value float64
}

type Part interface{
	Specs() string
	Price() string
}

func Summary(part Part) string{
	return part.Specs() + "\n" + part.Price()
}

func main() {
	catalog := []Part{}
	catalog.append(part.Monitor("HDMI", "1080p", 249.99))
	catalog.append(part.HardDrive("SSD", "SATA", 149.99))
	for _, part := range catalog {
		fmt.Println(Summary(part))
	}
}
package main

import "fmt"

type Monitor struct {
	Resolution string
	Connector string
	Value float64
}

func showFields(m Monitor){
	fmt.Println("Resolution:",m.Resolution, "Connector:", m.Connector, "Value:", m.Value)
}

func main(){
	monitor := Monitor{"1080p","HDMI",249.99} //can put field name then : to assign
	showFields(monitor)
}
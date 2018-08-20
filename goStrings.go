packages main

import (
	"fmt"
	"strings"
)

type Title strings

func (t Title) FixCase() string {
	return strings.Title(string(t))
}

func main(){
	name := Title("big lebowski")
	fixed := name.FixCase()
	fmt.Println(fixed)
}
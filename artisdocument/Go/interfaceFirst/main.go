package main

import (
	"fmt"
	"interface1/data"
)

func main() {
	n := data.NewName("kenji")
	fmt.Println(n.Value)
}



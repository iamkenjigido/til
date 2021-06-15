package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {
	b := []byte(`{"name":"hoge","age":30}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", p)
}

package main

import "fmt"

type Animal struct {
	crying string
}


func (a Animal) cry() {
	fmt.Println(a.crying)
}

func main() {
	elephant := Animal{crying: "paoon"}
	elephant.cry()

	monkey := Animal{crying: "ukiki"}
	monkey.cry()

	dog := Animal{crying: "wanwant"}
	dog.cry()

	cat := Animal{crying: "nya-nya"}
	cat.cry()
}

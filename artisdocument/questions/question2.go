package main

import "fmt"

type Human struct {
	name string
	age int
}

func (h Human) introduce() {
	fmt.Printf("私の名前は%vで%v歳です。\n", h.name, h.age)
}

func (h *Human) takeAge() {
	h.age ++
}

func (h *Human) takeAgeN(n int) {
	h.age += n
}

func main() {
	shimojima := Human{name: "Shimojima", age: 20}
	tamura := Human{name: "Tamura", age: 36}

	shimojima.introduce()
	shimojima.takeAge()
	shimojima.introduce()

	fmt.Println()

	tamura.introduce()
	tamura.takeAge()
	tamura.introduce()

	fmt.Println()

	shimojima.takeAgeN(10)
	shimojima.introduce()

	tamura.takeAgeN(63)
	tamura.introduce()
}

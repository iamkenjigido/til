package main

import (
	"fmt"
	"sort"
)

type Card struct {
	num []int
}

type Holder struct {
	card Card
}

type Reader struct {
	holder Holder
}

// 100枚のカードを作る
func (c *Card) makeCard() {
	for i := 100; i >= 1; i-- {
		c.num = append(c.num, i)
	}
}

// holderがソートしてcardを持つ
func (h *Holder) sort() {
	sort.Ints(h.card.num)
}

func (r *Reader) read() {
	for _, n := range r.holder.card.num {
		fizzBuzz(n)
	}
}

func fizzBuzz(number int) {
	if number % 15 == 0 {
		fmt.Println("FizzBuzz")
	} else if number % 5 == 0 {
		fmt.Println("Buzz")
	} else if number % 3 == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println(number)
	}
}


func main() {
	cards := Card{}
	cards.makeCard()
	holder := Holder{card: cards}
	holder.sort()
	reader := Reader{holder: holder}
	reader.read()
}
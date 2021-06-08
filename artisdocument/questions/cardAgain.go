package main

import (
	"fmt"
	"sort"

	//"sort"
)

type Card struct {
	num int
}

type Holder struct {
	card []Card
}

type Reader struct {
	holder Holder
}

// 100枚のカードを作る
func (c *Card) makeCard(num int) {
	c.num = num
}

//holderがカードを持つ
func (h *Holder) addCard(cards []Card) {
	for _, card := range cards {
		h.card = append(h.card, card)
	}
}
// ------------------------------- addCardから

//holderがソートしてcardを持つ
func (h *Holder) sort() {

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
	card := Card{}
	var cards []Card
	for i := 0; i < 100; i++ {
		card.makeCard(i)
		cards = append(cards, card)
	}
	holder := Holder{}
	holder.addCard(cards)
	holder.sort()
	//fmt.Println(holder.card)
	//reader := Reader{holder: holder}
	//reader.read()
}
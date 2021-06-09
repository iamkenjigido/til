package main

import (
	"fmt"
	"sort"
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
func (h Holder) addCard(cards []Card) {
	for _, card := range cards {
		h.card = append(h.card, card)
	}
}

//holderがソートしてcardを持つ
func (h *Holder) sort(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].num < cards[j].num
	})
	return cards
}

//func (r *Reader) read() {
//	for _, n := range r.holder.card {
//		fizzBuzz(n)
//	}
//}
func (r Reader) read(cards []Card) {
	for _, card := range cards {
		fizzBuzz(card)
	}
}

func fizzBuzz(number Card) {
	if number.num%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if number.num%5 == 0 {
		fmt.Println("Buzz")
	} else if number.num%3 == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println(number.num)
	}
}

func main() {

	card := Card{}
	var cards []Card
	for i := 100; i > 0; i-- {
		card.makeCard(i)
		cards = append(cards, card)
	}
	holder := Holder{}
	holder.addCard(cards)
	holder.sort(cards)
	fmt.Println(cards)
	reader := Reader{}
	reader.read(cards)
}

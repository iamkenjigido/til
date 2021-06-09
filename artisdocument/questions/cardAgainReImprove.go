package main

import (
	"fmt"
	"sort"
)

type Number struct {
	value int
}
type Card struct {
	number Number
}
type Holder struct {
	card []Card
}
type Reader struct {
	holder Holder
}

func makeNumber(i int) Number {
	n := Number{}
	n.value = i
	return n
}
func makeCard(n Number) Card {
	c := Card{}
	c.number = n
	return c
}
func (h *Holder) addCard(c Card) {
	h.card = append(h.card, c)
}
func (h *Holder) sort() {
	sort.Slice(h.card, func(i, j int) bool {
		return h.card[i].number.value < h.card[j].number.value
	})
}
func (r Reader) read() {
	fizzBuzzNum := 15
	BuzzNum := 5
	FizzNum := 3
	for _, v := range r.holder.card {
		switch {
		case v.number.value%fizzBuzzNum == 0:
			fmt.Println("FizzBuzz")
		case v.number.value%BuzzNum == 0:
			fmt.Println("Buzz")
		case v.number.value%FizzNum == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(v.number.value)
		}
	}
}

func main() {
	h := Holder{}
	for i := 100; i > 0; i-- {
		n := makeNumber(i)
		c := makeCard(n)
		h.addCard(c)
	}
	h.sort()
	r := Reader{}
	r.holder = h
	r.read()
}

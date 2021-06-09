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

func NewNumber(i int) Number {
	n := Number{}
	n.value = i
	return n
}

func NewCard(n Number) Card {
	c := Card{}
	c.number = n
	return c
}

func (h *Holder) AddCard(c Card) {
	h.card = append(h.card, c)
}

func (r *Reader) Sort() {
	sort.Slice(r.holder.card, func(i, j int) bool {
		return r.holder.card[i].number.value < r.holder.card[j].number.value
	})
}

func (r Reader) Print() {
	for _, v := range r.holder.card {
		switch {
		case v.number.value%15 == 0:
			fmt.Println("FizzBuzz")
		case v.number.value%5 == 0:
			fmt.Println("Buzz")
		case v.number.value%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(v.number.value)
		}
	}
}

func main() {
	h := Holder{}
	for i := 100; i > 0; i-- {
		n := NewNumber(i)
		c := NewCard(n)
		h.AddCard(c)
	}
	r := Reader{}
	r.holder = h
	r.Sort()
	r.Print()
}

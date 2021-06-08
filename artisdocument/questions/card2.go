package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"sort"
)

type Human struct {
	card []Card
}
type Card struct {
	num int
}

func makeCard() []int {
	var cards []int
	for i := 0; i < 100; i++ {
		c := Card{num: i + 1}
		cards = append(cards, c.num)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))
	log.Println(cards)
	return cards
}

func hold(card []int) {
	sort.Ints(card)
}

func read(card []int) {
	for _, c := range card {
		fizzBuzz(c)
	}
}

func fizzBuzz(number int) {
	if number%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if number%5 == 0 {
		fmt.Println("Buzz")
	} else if number%3 == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println(number)
	}
}

func main() {
	cards := makeCard()

	hold(cards)
	read(cards)

	//0から99のランダムな数値を表示する
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

}

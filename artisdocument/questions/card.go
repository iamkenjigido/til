package main

import "fmt"

type cardNum int

type card struct {
	num cardNum
}

type cardHolder struct {
	sortCard cardNum
}

func (c card) makeCard() {
	for i := 1; i <= 100; i++ {
		c.num = cardNum(i)
	}
}

func (c cardHolder) sort() {

}

func FizzBuzz(number int) {
	if number % 15 == 0 {
		fmt.Println("FizzBuzz")
	} else if number % 5 == 0 {
		fmt.Println("Buzz")
	} else if number % 3 == 0 {
		fmt.Println("Fizz")
	}
}

func main() {




	FizzBuzz(45)
}

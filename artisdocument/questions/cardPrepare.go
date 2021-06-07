package main

import (
	"fmt"
	"sort"
)

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
	var numbers []int
	for i := 1; i <= 100; i++ {
		numbers = append(numbers, i)
	}
	sort.Sort(sort.IntSlice(numbers))
	//fmt.Println(numbers)

	for _, number := range numbers {
		FizzBuzz(number)
	}

}

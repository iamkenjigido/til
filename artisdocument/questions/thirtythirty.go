package main

import "fmt"

func main() {
	//var numbers []int
	var sum int

	for i := 1; i <= 30; i++ {
		for j := 1; j <= 30; j++ {
			//numbers = append(numbers, i * j)
			fmt.Print(i * j, "\t")
			sum += i * j
		}
		fmt.Println(sum)
	}

	//fmt.Println(numbers)
}
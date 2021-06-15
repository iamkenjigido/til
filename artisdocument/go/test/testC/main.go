package main

import (
	"fmt"
	"sort"
)

func Input(border, reject int, score []int) int {
	var passed []int
	for _, v := range score {
		if border <= v {
			passed = append(passed, v)
		}
	}

	Sort(passed)
	result := Offer(reject, passed)
	return result
}

func Sort(passed []int) []int {
	sort.Sort(sort.IntSlice(passed))
	return passed
}

func Offer(reject int, passed []int) int {
	if len(passed) >= reject {
		passed = Reverse(passed)
		return CalculateOffer(passed, reject)
	}
	return 0
}

func Reverse(passed []int) []int {
	for i := 0; i < len(passed)/2; i++ {
		passed[i], passed[len(passed)-i-1] = passed[len(passed)-i-1], passed[i]
	}
	return passed
}

func CalculateOffer(passed []int, reject int) int {
	passed = passed[reject:]
	return len(passed)
}

func main() {
	border := 3974
	reject := 0
	score := []int{2049, 4690, 6867, 3414, 460}
	fmt.Println(Input(border, reject, score))

	border = 7879
	reject = 5
	score2 := []int{9270, 6818, 2995, 3354, 5289, 7837, 6486, 9371, 1338, 1395}
	fmt.Println(Input(border, reject, score2))

}

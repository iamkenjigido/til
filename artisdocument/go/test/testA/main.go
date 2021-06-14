package main

import (
	"fmt"
	"strings"
)

func CountWords(s string) map[string]int {
	slice := strings.Split(s, " ")
	var textMap = map[string]int{}
	for _, v := range slice {
		textMap[v] += 1
	}
	return textMap
}

func Print(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	var text1 = "red green blue blue green blue"
	textMap1 := CountWords(text1)
	fmt.Println(textMap1)
	Print(textMap1)

	var text2 = "Apple Apricot Orange Cherry Apple Orange Cherry Orange"
	textMap2 := CountWords(text2)
	Print(textMap2)

}

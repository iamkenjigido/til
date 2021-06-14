package main

import (
	"testing"
)

func TestTableCountWords(t *testing.T) {

	for _, v := range []struct {
		input    string
		expected map[string]int
	}{
		{input: "red green blue blue green blue", expected: map[string]int{
			"red":   1,
			"green": 2,
			"blue":  3,
		}},
		{input: "Apple Apricot Orange Cherry Apple Orange Cherry Orange", expected: map[string]int{
			"Cherry":  2,
			"Apple":   2,
			"Apricot": 1,
			"Orange":  3,
		}},
	} {
		output := CountWords(v.input)
		for key, _ := range output {
			if output[key] != v.expected[key] {
				t.Errorf("Test Faild: キーとバリューの対応が正しくないです")
			}
		}
	}
}

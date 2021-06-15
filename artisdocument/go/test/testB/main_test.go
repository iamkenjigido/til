package main

import (
	"testing"
)

func TestTableEvaluate(t *testing.T) {
	for _, v := range []struct {
		inputCorrect int
		inputYour    int
		expected     string
	}{
		{123456, 123456, "first"},
		{123456, 123455, "adjacent"},
		{123456, 123457, "adjacent"},
		{123456, 233456, "second"},
		{142358, 167358, "third"},
		{142358, 195283, "blank"},
	} {
		output := Evaluate(v.inputCorrect, v.inputYour)
		if output != v.expected {
			t.Errorf("Test failed: 当選番号と購入した番号の評価が間違っています")
		}
	}
}

func TestTableValidate(t *testing.T) {
	for _, v := range []struct {
		inputCorrect int
		inputYour int
	}{
		{1,1},
		{11,11},
		{111, 111},
		{1111, 1111},
		{11111, 11111},
		{1111111, 1111111},
		{11111111, 11111111},
	}{
		output := Validate(v.inputCorrect, v.inputYour)
		if output == nil {
			t.Errorf("Test faild: 戻り値が間違っています")
		}
	}
}

func TestTableValidate3(t *testing.T) {
	for _, v := range []struct {
		inputCorrect int
		inputYour int
		expected error
	}{
		{111111, 111111, nil},
	}{
		output := Validate(v.inputCorrect, v.inputYour)
		if output != v.expected {
			t.Errorf("Test faild: your input was %v and %v, but return %v", v.inputCorrect, v.inputYour, v.expected)
		}
	}
}
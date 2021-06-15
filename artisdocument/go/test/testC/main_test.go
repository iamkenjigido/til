package main

import (
	"testing"
)

func TestTableInput(t *testing.T) {
	for _, v := range []struct {
		border   int
		reject   int
		score    []int
		expected int
	}{
		{3974, 0, []int{2049, 4690, 6867, 3414, 460}, 2},
		{7879, 5, []int{9270, 6818, 2995, 3354, 5289, 7837, 6486, 9371, 1338, 1395}, 0},
		{5, 5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1},
		{4, 5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2},
	} {
		output := Input(v.border, v.reject, v.score)
		if output != v.expected {
			t.Errorf("Test faild: 戻り値が間違っています。")
		}
	}
}

func TestTableSort(t *testing.T) {
	for _, v := range []struct {
		passed   []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
	} {
		output := Sort(v.passed)
		for i := 0; i < len(output); i++ {
			if output[i] != v.expected[i] {
				t.Errorf("Test failed: 正しく昇順に並べ替えられていないです")
			}
		}
	}
}

func TestTableReverse(t *testing.T) {
	for _, v := range []struct {
		passed   []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
	} {
		output := Reverse(v.passed)
		for i := 0; i < len(output); i++ {
			if output[i] != v.expected[i] {
				t.Errorf("Test faild: 正しく降順に並べ替えられていないです")
			}
		}
	}
}

func TestTableCalculateOffer(t *testing.T) {
	for _, v := range []struct {
		passed   []int
		reject   int
		expected int
	}{
		{[]int{3, 2, 1}, 2, 1},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 5, 4},
	} {
		output := CalculateOffer(v.passed, v.reject)
		if output != v.expected {
			t.Errorf("Test faild: オファーを出す人数の計算が違います")
		}
	}
}

func TestTableOffer(t *testing.T) {
	for _, v := range []struct {
		reject   int
		passed   []int
		expected int
	}{
		{1, []int{3, 2, 1}, 2},
		{2, []int{4, 3, 2, 1}, 2},
		{3, []int{5, 4, 3, 2, 1}, 2},
		{5, []int{3, 2, 1}, 0},
		{6, []int{6, 5, 4, 3, 2, 1}, 0},
	} {
		output := Offer(v.reject, v.passed)
		if output != v.expected {
			t.Errorf("Test faild: オファーが間違っています")
		}
	}
}

package main

import "fmt"

type Common interface {
	countUp() Common
}

type hoge struct {
	num int
}

func (h hoge) countUp() Common {
	hh := hoge{}
	hh.num = h.num + 1
	return hh
}

type fuga struct {
	num int
}

func (f fuga) countUp() Common {
	ff := fuga{}
	ff.num = ff.num + 2
	return ff
}

type countUpper interface {
	Common
}

func main() {
	var buff []countUpper
	h1 := hoge{}
	h2 := fuga{}
	buff = append(buff, h1, h2)
	for _, v := range buff {
		vv := v.countUp()
		fmt.Println(vv)
	}
}

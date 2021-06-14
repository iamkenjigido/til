package main

import (
	"fmt"
	"github.com/kenji/til/artisdocument/go/interfaceFirstB1/structure"
)

var s []interface{}

type Namer interface {
	Name() string
}

func main() {
	s = append(s, structure.Hoge{}, structure.Fuga{}, structure.Moge{})

	for _, v := range s {
		switch value := v.(type) {
		case structure.Hoge:
			fmt.Println(value.HogeName("hoge"))
		case structure.Fuga:
			fmt.Println(value.FugaName("fuga"))
		case structure.Moge:
			fmt.Println(value.MogeName("moge"))
		}
	}
}


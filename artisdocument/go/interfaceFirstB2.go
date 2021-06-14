package main

import "fmt"

type Hoge struct {
	name string
}
type Fuga struct {
	name string
}
type Moge struct {
	name string
}
type Namer interface {
	Name() string
}

func (h Hoge) HogeName() string {
	return h.name
}
func (f Fuga) FugaName() string {
	return f.name
}
func (m Moge) MogeName() string {
	return m.name
}

func (h Hoge) Name() string {
	return h.name
}
func (f Fuga) Name() string {
	return f.name
}
func (m Moge) Name() string {
	return m.name
}

func main() {
	h := Hoge{"hoge"}
	f := Fuga{"fuga"}
	m := Moge{"moge"}

	s := []interface{}{h, f, m}

	for _, v := range s {
		switch value := v.(type) {
		case Hoge:
			fmt.Println(value.HogeName())
		case Fuga:
			fmt.Println(value.FugaName())
		case Moge:
			fmt.Println(value.MogeName())
		}
	}

	for _, v := range s {
		// Namerインターフェイスに型アサーション
		if value, ok := v.(Namer); ok {
			fmt.Println(value.Name())
		}
	}

}


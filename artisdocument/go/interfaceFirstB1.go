package main

import (
	"fmt"
)

type Hoge struct {
	name string
}
type Fuga struct {
	name string
}
type Moge struct {
	name string
}

func (h *Hoge) HogeName(name string) string {
	h.name = name
	return h.name
}
func (f *Fuga) FugaName(name string) string {
	f.name = name
	return f.name
}
func (m *Moge) MogeName(name string) string {
	m.name = name
	return m.name
}

type Namer interface {
	Name(name string) string
}

func (h Hoge) Name(name string) string {
	h.name = name
	return h.name
}
func (f Fuga) Name(name string) string {
	f.name = name
	return f.name
}
func (m Moge) Name(name string) string {
	m.name = name
	return m.name
}

func main() {
	var s []interface{}
	s = append(s, Hoge{}, Fuga{}, Moge{})
	for _, v := range s {
		switch value := v.(type) {
		case Hoge:
			fmt.Println(value.HogeName("hoge"))
		case Fuga:
			fmt.Println(value.FugaName("fuga"))
		case Moge:
			fmt.Println(value.MogeName("moge"))
		}
	}

	for _, v := range s {
		if value, ok := v.(Hoge); ok {
			fmt.Println(value.Name("hogehoge"))
		}
		if value, ok := v.(Fuga); ok {
			fmt.Println(value.Name("fugafuga"))
		}
		if value, ok := v.(Moge); ok {
			fmt.Println(value.Name("mogemoge"))
		}
	}

}


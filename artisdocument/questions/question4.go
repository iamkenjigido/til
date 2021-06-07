package main

import "fmt"

type Laundry struct {
	WashingMachine
	DryingMachine
}
type WashingMachine struct {
	kind string
	power string
}
type DryingMachine struct {
	kind string
	power string
}

type Clothes struct {
	kind string
	color string
}

func (w WashingMachine) wash(clothes Clothes)  {
	fmt.Printf("洗濯機の%v式を使ってパワー%vで%v色の%vを洗浄します\n", w.kind, w.power, clothes.color, clothes.kind)
}

func (d DryingMachine) dry(clothes Clothes) {
	fmt.Printf("乾燥器の%v式を使ってパワー%vで%v色の%vを乾かします\n", d.kind, d.power, clothes.color, clothes.kind)
}

func main() {
	strongLaundry := Laundry{WashingMachine{kind: "latest", power: "max"}, DryingMachine{kind: "latest", power: "max"}}
	clothes := Clothes{kind: "T-shirt", color: "black"}

	strongLaundry.wash(clothes)
	strongLaundry.dry(clothes)

}

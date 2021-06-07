package main

import "fmt"

type Car struct {
	Manufacturer string
	kind string
	Tire []Tire
}

func (c Car) drive() {
	fmt.Printf("%v %vは快適な走行を提供します。\n", c.Manufacturer, c.kind)
}

type Tire struct {
	amount int
}

func (t *Tire) addTires(n int) {
	if t.amount + n > 4 {
		fmt.Println("もうタイヤは4本ついてます。")
		return
	}
	if n > 4 {
		fmt.Println("タイヤは４本までしか取り付けられません")
		return
	}
	t.amount += n
	fmt.Printf("タイヤは今%v本ついてます。\n", t.amount)
}

func (c Car) now(location string) {
	fmt.Printf("現在地は%vです\n", location)
}
func (c Car) goStraight(m int)  {
	fmt.Printf("まっすぐ%vm進みました\n", m)
}
func (c Car) turnRight() {
	fmt.Println("右に曲がりました")
}
func (c Car) turnLeft()  {
	fmt.Println("左に曲がりました")
}

func main() {
	XTrail := Car{Manufacturer: "Nissan", kind: "X-Trail"}
	YarisCross := Car{Manufacturer: "Toyota", kind: "Yaris-Cross"}
	StepWagon := Car{Manufacturer: "Honda", kind: "Step-Wagon"}


	XTrail.drive()
	YarisCross.drive()
	StepWagon.drive()

	XTrail.Tire.addTires(4)

	XTrail.goStraight(100)
	XTrail.turnLeft()
}

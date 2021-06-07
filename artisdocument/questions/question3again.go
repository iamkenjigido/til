package main

import "fmt"

type Car struct {
	// 車です
	name Name
	tire []Tire
	location Location
}

type Name struct {
	//車の名前です
	value string
}

type Tire struct {
	// タイヤです
	manufacturer string
}

type Location struct {
	// 現在地です
	value int
}

// タイヤを付けます
func (c *Car) addTire(tires []Tire) error {
	for _, t := range tires {
		c.tire = append(c.tire, t)
	}
	fmt.Println("タイヤが付けられました")

	if !c.validateTireAmount() {
		return fmt.Errorf("タイヤの数が4本ではありません。")
	}

	fmt.Printf("タイヤの数は%v本です\n", len(c.tire))

	return nil
}

// タイヤの数を検証します
func (c *Car) validateTireAmount() bool {
	return len(c.tire) == 4
}

// 現在地を設定します。

//初期値は5とします。
func (c *Car) setInitialLocation() {
	c.location = Location{value: 5}
}

//右に曲がります
//現在地に1を足します
func (c *Car) turnRight() {
	c.location.value += 1
}

//左に曲がります
//現在地から1を引きます
func (c *Car) turnLeft() {
	c.location.value -= 1
}


func main() {
	//car := Car{name: Name{value: "ヤリスクロス"}}
	car := Car{}
	car.name.value = "ヤリスクロス"

	// タイヤを４本作成
	tireA := Tire{manufacturer: "Bridgestone"}
	tireB := Tire{manufacturer: "Bridgestone"}
	tireC := Tire{manufacturer: "Yokohama"}
	tireD := Tire{manufacturer: "Yokohama"}

	tires := []Tire{tireA, tireB, tireC, tireD}
	//fmt.Printf("%+v", tires)
	//car.addTire(tires)
	//fmt.Println(tires)
	if err := car.addTire(tires) ; err != nil {
		panic(err)
	}

	car.setInitialLocation()
	car.turnRight()
	car.turnLeft()

	fmt.Printf("車名: %s 位置: %d", car.name.value, car.location.value)

}


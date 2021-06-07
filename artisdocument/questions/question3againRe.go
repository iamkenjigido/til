package main

import "fmt"

type Car struct {
	name string
	tire []Tire
	location Location
}

type Tire struct {
	manufacture string
}

type Location struct {
	x int
	y int
}


// タイヤの装着
func (c *Car) addTires(tires []Tire) error {
	for _, t := range tires {
		c.tire = append(c.tire, t)
	}
	if !c.validatesAddTires() {
		return fmt.Errorf("タイヤの数が4本ではありません")
	}
	//c.validatesAddTires()
	fmt.Printf("タイヤを%d本装着しました \n", len(c.tire))

	return nil
}

// タイヤの検証
func (c *Car) validatesAddTires() bool {
	return len(c.tire) == 4
}


// 車の移動
func (c *Car) setInitialLocation() {
	c.location = Location{x: 0, y: 0}
}
func (c *Car) turnRight(dis int) {
	c.location.x += dis
}
func (c *Car) turnLeft(dis int) {
	c.location.x -= dis
}
func (c *Car) goStraight(dis int) {
	c.location.y += dis
}
func (c *Car) backStraight(dis int) {
	c.location.y -= dis
}

func (c *Car) move() {
	c.setInitialLocation()

}


func main() {
	car1 := Car{name: "日産エクストレイル"}

	tire1 := Tire{manufacture: "bridgestone"}
	tire2 := Tire{manufacture: "bridgestone"}
	tire3 := Tire{manufacture: "Yokohama"}
	tire4 := Tire{manufacture: "Yokohama"}

	tires := []Tire{tire1, tire2, tire3, tire4}

	if err := car1.addTires(tires); err != nil {
		panic(err)
	}

	car1.setInitialLocation()
	car1.turnRight(4)
	car1.turnLeft(2)
	car1.goStraight(7)
	car1.backStraight(2)
	fmt.Printf("車名: %v 現在地: (%d, %d) \n", car1.name, car1.location.x, car1.location.y)

}


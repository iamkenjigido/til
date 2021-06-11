package main

import "fmt"

type Name struct {
	value string
}
type Score struct {
	value int
}
type Child struct {
	name  Name
	score Score
}
type Class struct {
	name  string
	child []Child
}
type School struct {
	class []Class
}

func MakeName(name string) Name {
	n := Name{}
	n.value = name
	return n
}

func MakeScore(score int) Score {
	s := Score{}
	s.value = score
	return s
}

func (c Child) SetChildDate(data map[string]int) []Child {
	var children []Child
	for k, v := range data {
		c.ChildData(k, v)
		children = append(children, c)
	}
	return children
}

func (c Class) SetClassData(childrenMap map[string][]Child) []Class {
	var classes []Class
	for k, v := range childrenMap {
		c.ClassData(k, v)
		classes = append(classes, c)
	}
	return classes
}

func (c *Child) ChildData(name string, score int) {
	c.name = MakeName(name)
	c.score = MakeScore(score)
}

func (c *Class) ClassData(name string, child []Child) {
	c.name = name
	c.child = child
}

func (s School) Announce() {
	for _, c := range s.class {
		fmt.Printf("クラス名: %s \n", c.name)
		c.Print()
		//fmt.Println()
	}
}

func (c Class) Print() {
	for _, child := range c.child {
		fmt.Printf("名前: %sさん 点数: %d点 \n", child.name.value, child.score.value)
	}
	avg := c.Average()
	fmt.Printf("平均点は%d点です \n\n", avg)
}

func (c Class) Average() int {
	var sum int
	for _, child := range c.child {
		sum += child.score.value
	}
	avg := sum / len(c.child)
	return avg
}

func main() {
	child := Child{}
	var flowerData = map[string]int{
		"A": 80,
		"B": 90,
		"C": 100,
	}
	childrenFlower := child.SetChildDate(flowerData)

	var moonDate = map[string]int{
		"D": 60,
		"E": 100,
		"F": 30,
	}
	childrenMoon := child.SetChildDate(moonDate)

	var snowData = map[string]int{
		"G": 50,
		"H": 90,
	}
	childrenSnow := child.SetChildDate(snowData)

	class := Class{}
	var childrenMap = map[string][]Child{
		"花組": childrenFlower,
		"月組": childrenMoon,
		"雪組": childrenSnow,
	}
	classes := class.SetClassData(childrenMap)

	school := School{}
	school.class = classes
	school.Announce()

}

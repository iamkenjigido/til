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
	name string
	child []Child
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

func (c *Child) AddFeatureChild(name string, score int) /* map[Name]Score*/ {
	c.name, c.score = MakeName(name), MakeScore(score)
	//var childrenMap = map[Name]Score{c.name: c.score}
	//return childrenMap
}

func (c Class) Average() int {
	var sum int
	for _, child := range c.child {
		sum += child.score.value
	}
	avg := sum / len(c.child)
	return avg
}

func (c Class) Print() {
	fmt.Println(c.name)
	for _, child := range c.child {
		fmt.Printf("名前: %s 点数: %d \n", child.name.value, child.score.value)
	}
	avg := c.Average()
	fmt.Printf("平均点は%d点です \n", avg)
}

func main() {
	class := Class{}

	childA := Child{}
	childA.AddFeatureChild("A", 80)
	childB := Child{}
	childB.AddFeatureChild("B", 90)
	childC := Child{}
	childC.AddFeatureChild("C", 100)
	children := []Child{childA, childB, childC}

	class.name = "花組"
	class.child = children
	class.Print()
	//fmt.Println(class.name, class.child[0].name.value, class.child[0].score.value)
	//c := childA.AddFeatureChild("A", 80) 平均出すのに考えた
	//fmt.Println(c[Name{"A"}].value)


}

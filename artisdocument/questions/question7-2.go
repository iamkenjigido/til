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

func (c *Child) SetFeatureChild(name string, score int) {
	c.name, c.score = MakeName(name), MakeScore(score)
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
	for _, child := range c.child {
		fmt.Printf("名前: %sさん 点数: %d点 \n", child.name.value, child.score.value)
	}
	avg := c.Average()
	fmt.Printf("平均点は%d点です \n\n", avg)
}

func (c *Class) SetFeatureClass(name string, child []Child) {
	c.name = name
	c.child = child
}

func (s School) Announce() {
	for _, c := range s.class {
		fmt.Printf("クラス名: %s \n", c.name)
		c.Print()
	}
}

func main() {
	childA := Child{}
	childA.SetFeatureChild("A", 80)
	childB := Child{}
	childB.SetFeatureChild("B", 90)
	childC := Child{}
	childC.SetFeatureChild("C", 100)
	childrenFlower := []Child{childA, childB, childC}
	classFlower := Class{}
	classFlower.SetFeatureClass("花組", childrenFlower)

	childD := Child{}
	childD.SetFeatureChild("D", 60)
	childE := Child{}
	childE.SetFeatureChild("E", 100)
	childF := Child{}
	childF.SetFeatureChild("F", 30)
	childrenMoon := []Child{childD, childE, childF}
	classMoon := Class{}
	classMoon.SetFeatureClass("月組", childrenMoon)

	childG := Child{}
	childG.SetFeatureChild("G", 50)
	childH := Child{}
	childH.SetFeatureChild("H", 90)
	childrenSnow := []Child{childG, childH}
	classSnow := Class{}
	classSnow.SetFeatureClass("雪組", childrenSnow)

	classes := []Class{classFlower, classMoon, classSnow}
	school := School{}
	school.class = classes
	school.Announce()
}

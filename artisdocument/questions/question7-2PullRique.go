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
	childrenFlower := []Child{}
	for _, v := range []struct {
		name  string
		score int
	}{
		{name: "A", score: 80},
		{name: "B", score: 90},
		{name: "C", score: 100},
	} {
		c := Child{}
		c.SetFeatureChild(v.name, v.score)
		childrenFlower = append(childrenFlower, c)
	}
	classFlower := Class{}
	classFlower.SetFeatureClass("花組", childrenFlower)

	childrenMoon := []Child{}
	for _, v := range []struct {
		name string
		score int
	}{
		{name: "D", score: 60},
		{name: "E", score: 100},
		{name: "F", score: 30},
	}{
		c := Child{}
		c.SetFeatureChild(v.name, v.score)
		childrenMoon = append(childrenMoon, c)
	}
	classMoon := Class{}
	classMoon.SetFeatureClass("月組", childrenMoon)

	childrenSnow := []Child{}
	for _, v := range []struct {
		name string
		score int
	}{
		{name: "G", score: 50},
		{name: "H", score: 90},
	}{
		c := Child{}
		c.SetFeatureChild(v.name, v.score)
		childrenSnow = append(childrenSnow, c)
	}
	classSnow := Class{}
	classSnow.SetFeatureClass("雪組", childrenSnow)

	classes := []Class{classFlower, classMoon, classSnow}
	school := School{}
	school.class = classes
	school.Announce()
}

package main

import "fmt"
// テストの結果です
type TestResult struct {
	nameScore map[string]int
	average   int
}
// クラスです
type Class struct {
	name       string
	testResult TestResult
}
// テストの結果を作ります
func MakeTestResult(data map[string]int) TestResult {
	t := TestResult{}
	t.nameScore = data
	return t
}
// クラスを作ります
func MakeClass(name string, t TestResult) Class {
	c := Class{}
	c.name = name
	c.testResult = t
	return c
}
// 平均点を計算します
func (c *Class) CalcAverage() {
	var s int
	for _, v := range c.testResult.nameScore {
		s += v
	}
	avg := s / len(c.testResult.nameScore)
	c.testResult.average = avg
}
// 出力します
func (c Class) Print() {
	fmt.Printf("クラス：%s\n", c.name)
	for name, score := range c.testResult.nameScore {
		fmt.Printf("生徒名：%sさん 点数：%d\n", name, score)
	}
	fmt.Printf("平均点:%d\n\n", c.testResult.average)
}

func main() {
	// 花組です
	flower := map[string]int{
		"A": 80,
		"B": 90,
		"C": 100,
	}
	// 月組です
	moon := map[string]int{
		"D": 60,
		"E": 100,
		"F": 20,
	}
	// 雪組です
	snow := map[string]int{
		"G": 50,
		"H": 90,
	}

	f := MakeTestResult(flower)
	m := MakeTestResult(moon)
	s := MakeTestResult(snow)

	ff := MakeClass("花組", f)
	mm := MakeClass("月組", m)
	ss := MakeClass("雪組", s)

	for _, v := range []struct {
		class Class
	}{
		{ff},
		{mm},
		{ss},
	} {
		v.class.CalcAverage()
		v.class.Print()
	}

}

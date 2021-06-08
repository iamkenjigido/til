package main

import (
	"fmt"
	"time"
)

// BendingMachine   です
type BendingMachine struct {
	drink map[string]Drink
}

// Drink   です
type Drink struct {
	name  string
	price int
}

// Human    です
type Human struct {
	money int
}

// Drinkを作ります
func makeDrink(name string, price int) Drink {
	d := Drink{
		name:  name,
		price: price,
	}
	return d
}

// BendingMachineにdrinkをセットします
func (b *BendingMachine) setDrink(drink Drink) {
	b.drink[drink.name] = drink
}

// 人の持つお金を生成します
func (h Human) makeMoney(money int) int {
	h.money = money
	return h.money
}

// bendingMachine で購入します
func (b BendingMachine) buyDrink(election string, money int) {
	fmt.Println(election, "をお選びですね \n", b.drink[election].price, "円です")
	time.Sleep(3 * time.Second)
	if dif := money - b.drink[election].price; dif > 0 {
		fmt.Printf("お買い上げありがとうございました！おつりは%d円です。", dif)
	} else if dif == 0 {
		fmt.Println("ちょうどお預かりします。お買い上げありがとうございました！")
	} else {
		fmt.Printf("お金が%d円足りませんでしたので買うことはできませんでした。", -dif)
	}
}



func main() {
	drink1 := makeDrink("コカ・コーラ", 150)
	drink2 := makeDrink("ファンタグレープ", 140)
	drink3 := makeDrink("コーヒー", 120)
	drink4 := makeDrink("ナタデココ", 130)
	drinks := []Drink{drink1, drink2, drink3, drink4}
	for _, d := range drinks {
		fmt.Printf("%s:%d円  ", d.name, d.price)
	}

	bendingMachine := BendingMachine{map[string]Drink{}}
	for _, d := range drinks {
		bendingMachine.setDrink(d)
	}

	man := Human{}

	fmt.Println("\nお金を入れてください")
	var putIn int
	putIn, err := fmt.Scan(&putIn)
	if err != nil {
		return
	}
	use := man.makeMoney(putIn)

	fmt.Println("買うドリンクを教えてください")
	var election string
	fmt.Scan(&election)
	bendingMachine.buyDrink(election, use)
}

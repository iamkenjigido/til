package main

import "fmt"

type yen int
type drink string

// 自販機の構造体
type BendingMachine struct {
	kind drink
	price yen
}

// 支払うお金の構造体
type Payment struct {
	price yen
}

// 入れられたお金と飲み物の値段を比較する
func (b BendingMachine) evaluatePrice(putMoney yen) string {
	if b.price <= putMoney {
		fmt.Printf("%v円を入れたら ", putMoney)
		return "ok"
	} else {
		fmt.Printf("%v円を入れたら ", putMoney)
		return "NG"
	}
}

// 商品を出す
func (b BendingMachine) provideDrink(consequence string) {
	if consequence == "ok" {
		fmt.Printf("%vが買えました", b.kind)
	} else if consequence == "NG" {
		fmt.Printf("%v円では買えません、と怒られました", b.price)
	}
}

func main() {
	coke := BendingMachine{kind: "coke", price: 150}
	pay := Payment{price: 150}

	consequence := coke.evaluatePrice(pay.price)
	coke.provideDrink(consequence)

}

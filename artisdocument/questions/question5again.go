package main

import (
	"fmt"
)

type yen2 int
type drink2 string
type colors string

type BendingMachine2 struct {
	// 自販機の構造体
	color colors
	Juice
}

type Juice struct {
	//ジュースの構造体
	name drink2
	price yen2
}

type Payment2 struct {
	// 支払うお金の構造体
	price yen2
}

// 自販機を見つける
func (b BendingMachine2) findIt() {
	fmt.Printf("%v色の自販機をみつけた！", b.color)
}

// 商品を見て決める
func (b BendingMachine2) selectIt() drink2 {
	fmt.Printf("%v色の自販機の中に%vがあったので%vを買おうと思った。", b.color, b.name, b.name)
	selected
}

// お金を財布から取り出す
func (p Payment2) pickUp() yen2 {
	pickUpMoney := p.price
	return pickUpMoney
}

// お金を自販機に入れる
func (b *BendingMachine2) putIn(pickUpMonet yen2) {
	fmt.Printf("ガシャンガシャン、%v円が入りました。", pickUpMonet)
}

// 選んだ飲み物のボタンを押す
func (b BendingMachine2) pushButton() {
	
}

// 自販機の中でお金が足りているかの評価をする
func (b BendingMachine2) evaluate() {
	
}

// 足りていたら商品を出す
func (b BendingMachine2) putOut() {

}

// 「ありがとうございます」と表示する
func (b BendingMachine2) thank() {

}

func main() {
	redMachine := BendingMachine2{
		color: "red",
	}
	fanta := Juice{
		name: "ファンタグレープ",
		price: 150,
	}
	myPay1 := Payment2{
		price: 150,
	}

	redMachine.findIt()
	redMachine.selectIt()
	pickUpMoney := myPay1.pickUp()
	redMachine.putIn(pickUpMoney)
	redMachine.pushButton()




}

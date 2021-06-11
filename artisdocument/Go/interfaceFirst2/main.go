package main

import (
	"fmt"
	"github.com/kenji/til/artisdocument/Go/interfaceFirst2/customer"
	"github.com/kenji/til/artisdocument/Go/interfaceFirst2/leader"
	"github.com/kenji/til/artisdocument/Go/interfaceFirst2/manager"
	"github.com/kenji/til/artisdocument/Go/interfaceFirst2/part"
)

type Registerer interface {
	Register()
}

//func Payment(registerer Registerer, customer []string) {
//	for _, v := range customer {
//		registerer.Register()
//		fmt.Printf("%sのお会計が終わりました\n\n", v)
//	}
//}

func Payment(registerer Registerer, customer string)  {
	fmt.Printf("%sのお会計が終わりました\t", customer)
	registerer.Register()
}

func main() {
	var im, il, ip Registerer
	im = manager.NewManager("山田", "店長")
	il = leader.NewLeader("鈴木", "リーダー")
	ip = part.NewPart("田村", "バイト")

	//Payment(im, customer.Customer)
	//Payment(il, customer.Customer)
	//Payment(ip, customer.Customer)
	iii := []Registerer{im, il, ip}

	for _, v := range customer.Customer {
		for _, i := range iii {
			Payment(i, v)
		}
	}

}

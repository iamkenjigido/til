package main

import (
	"github.com/kenji/til/artisdocument/Go/interfaceFirst2/leader"
	"github.com/kenji/til/artisdocument/Go/interfaceFirst2/manager"
	"github.com/kenji/til/artisdocument/Go/interfaceFirst2/part"
)

type Registerer interface {
	Register()
}

func Payment(registerer Registerer) {
	registerer.Register()
}

func main() {

	var im, il, ip Registerer
	im = manager.NewManager("山田", "店長")
	il = leader.NewLeader("鈴木", "リーダー")
	ip = part.NewPart("田村", "バイト")

	Payment(im)
	Payment(il)
	Payment(ip)

}

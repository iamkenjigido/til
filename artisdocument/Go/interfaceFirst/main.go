package main

import (
	"fmt"
	"github.com/kenji/til/artisdocument/Go/interfaceFirst/data"
	"github.com/kenji/til/artisdocument/Go/interfaceFirst/staff"
)

type Registerer staff.Registerer
type Manager data.Manager
type Leader data.Leader
type Staff data.Staff

func (m Manager) Register() string {
	st := fmt.Sprintf("役職:%s、名前:%sがレジを打ちました\n", m.Person.Name.Value, m.Person.Post.Value)
	return st
}

func (l Leader) Register() string {
	st := fmt.Sprintf("役職:%s、名前:%sがレジを打ちました\n", l.Person.Name.Value, l.Person.Post.Value)
	return st
}

func (s Staff) Register() string {
	st := fmt.Sprintf("役職:%s、名前:%sがレジを打ちました\n", s.Person.Name.Value, s.Person.Post.Value)
	return st
}


func Payment(r string,name string, customer string) {
	fmt.Printf(r)
	fmt.Printf("%sさんが%sのお会計をしました\n", name, customer)
}

func main() {
	p := data.Person{}
	yamada := p.NewPerson("山田", "店長")
	suzuki := p.NewPerson("鈴木", "バイトリーダー")
	tamura := p.NewPerson("田村", "バイト")
	//fmt.Println(yamada, suzuki, tamura)

	customer := []string{"A", "B", "C"}

	var ry Registerer = Manager{Person: yamada}
	var rs Registerer = Manager{Person: suzuki}
	var rt Registerer = Manager{Person: tamura}

	a := ry.Register()
	b := rs.Register()
	c := rt.Register()

	Payment(a, yamada.Name.Value, customer[0])
	Payment(b, suzuki.Name.Value, customer[1])
	Payment(c, tamura.Name.Value, customer[2])
}

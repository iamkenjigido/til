package staff

import (
	"fmt"
	"github.com/kenji/til/artisdocument/Go/interfaceFirst/data"
)



type Registerer interface {
	Register() string
}

type Manager data.Manager
type Leader data.Leader
type Staff data.Staff

//func (m Manager) Register() string {
//	st := fmt.Sprintf("役職:%s、名前:%sがレジを打ちました\n", m.Person.Name.Value, m.Person.Post.Value)
//	return st
//}
//
//func (l Leader) Register() string {
//	st := fmt.Sprintf("役職:%s、名前:%sがレジを打ちました\n", l.Person.Name.Value, l.Person.Post.Value)
//	return st
//}
//
//func (s Staff) Register() string {
//	st := fmt.Sprintf("役職:%s、名前:%sがレジを打ちました\n", s.Person.Name.Value, s.Person.Post.Value)
//	return st
//}

//func Payment(r Registerer,name string, customer string) {
//	r.Register()
//	fmt.Printf("%sさんが%sのお会計をしました\n", name, customer)
//}


type Person data.Person

func (p Person) Register() string {
	s := fmt.Sprintf("役職：００、名前：００がレジを打ちました")
	return s
}

func (p Person) Payment(customer string) {
	fmt.Printf("%sさんが%sのお会計をしました。", p.Name.Value, customer)
}

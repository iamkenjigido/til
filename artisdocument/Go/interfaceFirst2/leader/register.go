package leader

import "fmt"

func (m Leader) Register() {
	fmt.Printf("役職：%s\t名前：%sがレジを打ちました\n", m.Post, m.Name)
}
package manager

import "fmt"

func (m Manager) Register() {
	fmt.Printf("役職：%s\t名前：%sがレジを打ちました\n\n", m.Post, m.Name)
}

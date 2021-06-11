package manager

type Manager struct {
	Name string
	Post string
}

func NewManager(name string, post string) Manager {
	m := Manager{}
	m.Name = name
	m.Post = post
	return m
}

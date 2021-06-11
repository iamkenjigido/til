package leader

type Leader struct {
	Name string
	Post string
}

func NewLeader(name string, post string) Leader {
	l := Leader{}
	l.Name = name
	l.Post = post
	return l
}

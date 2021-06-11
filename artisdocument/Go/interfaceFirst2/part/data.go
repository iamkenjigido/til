package part

type Part struct {
	Name string
	Post string
}

func NewPart(name string, post string) Part {
	p := Part{}
	p.Name = name
	p.Post = post
	return p
}

package data

type Manager struct {
	Person Person
}
type Leader struct {
	Person Person
}
type Staff struct {
	Person Person
}

type Person struct {
	Name Name
	Post Post
}

type Name struct {
	Value string
}

type Post struct {
	Value string
}

func NewName(name string) Name {
	n := Name{}
	n.Value = name
	return n
}

func NewPost(post string) Post {
	p := Post{}
	p.Value = post
	return p
}

func (p Person) NewPerson(name string, post string) Person {
	p.Name = NewName(name)
	p.Post = NewPost(post)
	return p
}


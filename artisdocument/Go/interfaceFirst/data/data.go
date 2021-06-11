package data

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
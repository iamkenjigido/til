package main

import (
	"encoding/json"
	"fmt"
)

// T構造体 埋め込み用
type T struct {
	title string
	description string
}

// jsonT 変換用構造体
type jsonT struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

// MarshalJSON T構造体用のカスタマイズ
func (t T) MarshalJSON() ([]byte, error) {
	// T型のフィールドがプライベートなのでjsonT型の構造体に値をコピーする
	jt := &jsonT{}
	jt.Title = t.title
	jt.Description = t.description
	b, err := json.Marshal(jt)
	if err != nil {
		return nil, err
	}
	return b, err
}

// Person 構造体
type Person struct {
	name string
	age int
	nicknames []string
	T
}

// MarshalJSON Person構造体用のカスタマイズ
func (p Person) MarshalJSON() ([]byte, error) {
	// 埋め込まれているTだけ一度jsonに変換する
	t, err := json.Marshal(&p.T)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(t, &data)
	if err != nil {
		return nil, err
	}

	data["name"] = p.name
	data["age"] = p.age
	data["nicknames"] = p.nicknames

	pj, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return pj, err
}

func main() {
	p := Person{
		name: "aaa",
		age: 25,
		nicknames: []string{"ccc", "ddd", "eee"},
		T: T{
			"fff",
			"ggggg",
		},
	}
	fmt.Printf("before struct: %+v\n", p)

}
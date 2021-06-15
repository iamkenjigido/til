package main

import (
	"fmt"
)

func Evaluate(i, ii int) string {
	switch {
	case i == ii:
		return fmt.Sprintf("first")
	case i-1 == ii || i+1 == ii:
		return fmt.Sprintf("adjacent")
	case i%10000 == ii%10000:
		return fmt.Sprintf("second")
	case i%1000 == ii%1000:
		return fmt.Sprintf("third")
	default:
		return fmt.Sprintf("blank")
	}
}

func Validate(i, j int) error {
	switch {
	case i < 100000 || 199999 < i:
		return fmt.Errorf("桁数が間違っています")
	case j < 100000 || 199999 < j:
		return fmt.Errorf("桁数が間違っています")
	default:
		return nil
	}
}

func main() {
	fmt.Println("当選番号と購入した宝くじの番号を入力してください")

	var result []string

	for i := 0; i < 3; i++ {
		var correct, your int
		fmt.Scan(&correct, &your)
		if err := Validate(correct, your); err != nil {
			panic(err)
		}
		result = append(result, Evaluate(correct, your))
	}
	for _, v := range result {
		fmt.Println(v)
	}

}

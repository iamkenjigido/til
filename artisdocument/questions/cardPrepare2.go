package main

type Currency string

// ドルとユーロ
const (
	USD = Currency("USD")
	EUR = Currency("EUR")
)

type Money struct {
	amount int
	currency Currency
}

// 通貨からMoneyを生成する
func (c Currency) New(a int) Money {
	return Money{amount: a, currency: c}
}

// 5ドルのお金はこうなる
//five := USD.New(5)

// 交換レート
type Rate int

// 通貨のペア（交換の方向も持つ）
type pair struct {
	from Currency
	to Currency
}

// 銀行（通貨のペア毎に交換レートを持つ）
type Bank struct {
	rates map[pair]Rate // pairがキーでRateがバリュー
}

//交換レートの追加
func (b Bank) AddRate(from Currency, to Currency, rate Rate) {
	b.rates[pair{from, to}] = rate
}

// 交換レートの参照
func (b Bank) RefRate(from Currency, to Currency) Rate {
	if rate, ok := b.rates[pair{from, to}]; ok {
		return rate
	}
	return Rate(1)
}

func newBank() Bank {
	return &bank{map[pair]Rate{}}
}
























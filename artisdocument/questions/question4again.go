package main

import (
	"fmt"
	"math"
)

// 洗濯乾燥機です
type Machine struct {
	washing Washing
	drying Drying
}
// 洗濯機部分です
type Washing struct {
	power int
	time int
}
// 乾燥機部分です
type Drying struct {
	power int
	time int
}
// 汚れた衣類です
type Clothing struct {
	kind string
	level int
	volume int
}


// 洗濯機の強さと洗う時間を決めます
func (m *Machine) setWashingStrength(strength, timer int) error {
	if strength > 10 {
		return fmt.Errorf("強さが10より大きくなっています")
	}
	m.washing.power = strength

	if timer > 10 {
		return fmt.Errorf("時間が10より大きくなっています")
	}
	m.washing.time = timer

	return nil
}

// 乾燥機の強さと乾燥時間を決めます
func (m *Machine) setDryingStrength(strength, timer int) error {
	if strength > 10 {
		return fmt.Errorf("強さが10より大きくなっています")
	}
	m.drying.power = strength

	if timer > 10 {
		return fmt.Errorf("時間が10より大きくなっています")
	}
	m.drying.time = timer

	return  nil
}

// 衣類の汚れ具合と量を決めます
func (c *Clothing) setClothingState(dirty, amount int) error {
	if dirty > 10 {
		return fmt.Errorf("levelが10より大きくなっています")
	}
	c.level = dirty * 10

	if amount > 10 {
		return fmt.Errorf("amountが10より大きくなっています")
	}
	c.volume = amount * 10

	return nil
}


// 洗濯機の洗濯の評価
func (m *Machine) evaluateWash(cloth Clothing)  {
	washTotalPower := m.washing.power * m.washing.time
	clothDirt := cloth.level
	if washTotalPower >= clothDirt {
		fmt.Printf("洗浄が完了しました \n")
	} else {
		lackTime := math.Ceil(float64((washTotalPower - clothDirt) / m.washing.power * -1))
		fmt.Printf("洗濯が完了するまで、もう%.f分お待ちください \n", lackTime)
	}
}

// 乾燥機の乾燥の評価
func (m *Machine) evaluateDry(cloth Clothing) {
	dryTotalPower := m.drying.power * m.drying.time
	clothAmount := cloth.volume
	if dryTotalPower >= clothAmount {
		fmt.Printf("乾燥が完了しました \n")
	} else {
		lackTime := math.Ceil(float64((dryTotalPower - clothAmount) / m.drying.power * -1))
		fmt.Printf("完全に乾くまで、もう%.f分お待ちください \n", lackTime)
	}
}


func main() {
	fmt.Println("kenjiMachineを使う")
	kenjiMachine := Machine{}
	if err := kenjiMachine.setWashingStrength(4, 3); err != nil {
		panic(err)
	}
	if err := kenjiMachine.setDryingStrength(8, 5); err != nil {
		panic(err)
	}


	cloth := Clothing{}
	if err := cloth.setClothingState(4, 6); err != nil {
		panic(err)
	}


	kenjiMachine.evaluateWash(cloth)
	kenjiMachine.evaluateDry(cloth)

	//fmt.Println("追加で洗濯したい時間を入力してください")
	//var lackWashTime int
	//fmt.Scan(&lackWashTime)
	//kenjiMachine.setWashingStrength(4, lackWashTime)
	//kenjiMachine.evaluateWash(cloth)

	fmt.Println("kyokoMachineを使う")
	kyokoMachine := Machine{}
	if err := kyokoMachine.setWashingStrength(7, 10); err != nil {
		panic(err)
	}
	if err := kyokoMachine.setDryingStrength(8, 10); err != nil {
		panic(err)
	}


	cloth = Clothing{}
	if err := cloth.setClothingState(7, 8); err != nil {
		panic(err)
	}


	kyokoMachine.evaluateWash(cloth)
	kyokoMachine.evaluateDry(cloth)

	fmt.Println("kaitoMachineを使う")
	kaitoMachine := Machine{}
	if err := kaitoMachine.setWashingStrength(3, 10); err != nil {
		panic(err)
	}
	if err := kaitoMachine.setDryingStrength(8, 10); err != nil {
		panic(err)
	}


	cloth = Clothing{}
	if err := cloth.setClothingState(4, 7); err != nil {
		panic(err)
	}


	kaitoMachine.evaluateWash(cloth)
	kaitoMachine.evaluateDry(cloth)

}
package main

import "fmt"

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



func main() {
	kenjiMachine := Machine{}
	if err := kenjiMachine.setWashingStrength(10, 10); err != nil {
		panic(err)
	}
	if err := kenjiMachine.setDryingStrength(8, 5); err != nil {
		panic(err)
	}

	cloth := Clothing{}

	if err := cloth.setClothingState(4, 6); err != nil {
		panic(err)
	}




}
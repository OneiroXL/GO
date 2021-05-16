package Computer

import (
	"Bridge/I"
	"fmt"
)


type LaptopComputer struct {
	Computer
}

/*
初始化
*/
func NewLaptopComputer(brand I.IBrand ) *LaptopComputer {
	laptopComputer := new(LaptopComputer)
	laptopComputer.Computer.Brand = brand
	return laptopComputer
}

/*

*/
func (this *LaptopComputer) CInfo() {
	this.Computer.CInfo()
	fmt.Println("笔记本")
}
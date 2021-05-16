package Computer

import (
	"Bridge/I"
	"fmt"
)

type DesktopComputer struct {
	Computer
}

/*
初始化
*/
func NewDesktopComputer(brand I.IBrand ) *DesktopComputer {
	desktopComputer := new(DesktopComputer)
	desktopComputer.Computer.Brand = brand
	return desktopComputer
}

/*

*/
func (this *DesktopComputer) CInfo() {
	this.Computer.CInfo()
	fmt.Println("台式机")
}
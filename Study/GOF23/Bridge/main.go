package main

import (
	"Bridge/C/Brand"
	"Bridge/C/Computer"
	"Bridge/I"
)

func main() {
	var desktopComputerA I.IComputer = Computer.NewDesktopComputer(new(Brand.Apple))
	desktopComputerA.CInfo()
	
	var desktopComputerL I.IComputer = Computer.NewDesktopComputer(new(Brand.Lenovo))
	desktopComputerL.CInfo()

	var laptopComputerA I.IComputer = Computer.NewLaptopComputer(new(Brand.Apple))
	laptopComputerA.CInfo()
	
	var laptopComputerL I.IComputer = Computer.NewLaptopComputer(new(Brand.Lenovo))
	laptopComputerL.CInfo()


}
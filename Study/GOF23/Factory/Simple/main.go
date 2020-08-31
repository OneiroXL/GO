package main

import (
	"Factory/Simple/I"
	"Factory/Simple/F"
)

func main(){
	var wuLin I.Car = F.GetCar("W")
	var tesla I.Car = F.GetCar("T")

	wuLin.GetName()
	tesla.GetName()
}
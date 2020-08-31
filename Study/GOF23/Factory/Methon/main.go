package main

import (
	"Factory/Methon/I"
	"Factory/Methon/F"
)


/*
工厂方法模式
*/

func main(){

	var wuLingFactory I.ICarFactory = new(F.WuLingFactory)
	var teslaFactory I.ICarFactory = new(F.TeslaFactory)

	var wuLin I.ICar = wuLingFactory.GetCar()
	var tesla I.ICar = teslaFactory.GetCar()
	wuLin.GetName()
	tesla.GetName()
}
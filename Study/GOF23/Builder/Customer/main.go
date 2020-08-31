package main

import (
	"Builder/Customer/C"
	"Builder/Customer/I"
)



func main(){
	var computerWorker I.IBuilder = C.NewComputerWorker()

	computerWorker.BuilderA("CPU").BuilderB("GPU").BuilderC("RAM").BuilderD("POWER")

	product := computerWorker.GetProduct();
	product.ToString()
}
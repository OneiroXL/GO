package main

import (
	"Builder/Director/Master"
	"Builder/Director/C"
	"Builder/Director/I"
)


func main(){

	master := new(Master.Master)
	var worker I.IBuilder = new(C.Worker)
	var computerWorker I.IBuilder = new(C.ComputerWorker)

	product := master.Build(worker)
	product.ToString()

	computerProduct := master.Build(computerWorker)
	computerProduct.ToString()
}
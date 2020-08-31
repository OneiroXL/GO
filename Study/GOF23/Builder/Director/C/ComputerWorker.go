package C



import (
	_"Builder/Director/I"
	"Builder/Director/C/Entity"
)

var (
	comProduct = new(Entity.Product)
)

type ComputerWorker struct{

}

func (this *ComputerWorker) BuilderA() {
	comProduct.SetA("CPU")
}

func (this *ComputerWorker) BuilderB() {
	comProduct.SetB("GPU")
}

func (this *ComputerWorker) BuilderC() {
	comProduct.SetC("RAM")
}

func (this *ComputerWorker) BuilderD() {
	comProduct.SetD("POWER")
}

func (this *ComputerWorker) GetProduct()  *Entity.Product {
	return comProduct
}

package C

import (
	"Builder/Customer/I"
	"Builder/Customer/C/Entity"
)


type ComputerWorker struct{
	product *Entity.Product
}

func NewComputerWorker() *ComputerWorker {
	return &ComputerWorker{
		product:new(Entity.Product),
	}
}

func (this *ComputerWorker) BuilderA(v string) I.IBuilder {
	this.product.SetA(v)
	return this
}

func (this *ComputerWorker) BuilderB(v string) I.IBuilder {
	this.product.SetB(v)
	return this
}

func (this *ComputerWorker) BuilderC(v string) I.IBuilder {
	this.product.SetC(v)
	return this
}

func (this *ComputerWorker) BuilderD(v string) I.IBuilder {
	this.product.SetD(v)
	return this
}

func (this *ComputerWorker) GetProduct() *Entity.Product {
	return this.product
}

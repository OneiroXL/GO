package C


import (
	"Builder/Customer/I"
	"Builder/Customer/C/Entity"
)

var (
	
)

type Worker struct{
	product *Entity.Product
}

func NewWorker() *Worker {
	return &Worker{
		product:new(Entity.Product),
	}
}

func (this *Worker) BuilderA(v string) I.IBuilder {
	this.product.SetA(v)
	return this
}

func (this *Worker) BuilderB(v string) I.IBuilder {
	this.product.SetB(v)
	return this
}

func (this *Worker) BuilderC(v string) I.IBuilder {
	this.product.SetC(v)
	return this
}

func (this *Worker) BuilderD(v string) I.IBuilder {
	this.product.SetD(v)
	return this
}

func (this *Worker) GetProduct() *Entity.Product {
	return this.product
}

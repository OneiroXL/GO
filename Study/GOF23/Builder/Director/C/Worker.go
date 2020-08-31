package C


import (
	_"Builder/Director/I"
	"Builder/Director/C/Entity"
)

var (
	product = new(Entity.Product)
)

type Worker struct{

}

func (this *Worker) BuilderA() {
	product.SetA("地基")
}

func (this *Worker) BuilderB() {
	product.SetB("钢筋")
}

func (this *Worker) BuilderC() {
	product.SetC("铺线")
}

func (this *Worker) BuilderD() {
	product.SetD("装修")
}

func (this *Worker) GetProduct()  *Entity.Product {
	return product
}

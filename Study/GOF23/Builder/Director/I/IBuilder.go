package I

import (
	"Builder/Director/C/Entity"
)

/*
抽象的建造者
*/
type IBuilder interface{
	BuilderA()
	BuilderB()
	BuilderC()
	BuilderD()
	GetProduct() *Entity.Product
}
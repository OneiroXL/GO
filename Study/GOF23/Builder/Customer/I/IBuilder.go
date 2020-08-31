package I

import (
	"Builder/Customer/C/Entity"
)

/*
抽象的建造者
*/
type IBuilder interface{
	BuilderA(v string) IBuilder
	BuilderB(v string) IBuilder
	BuilderC(v string) IBuilder
	BuilderD(v string) IBuilder
	GetProduct() *Entity.Product
}
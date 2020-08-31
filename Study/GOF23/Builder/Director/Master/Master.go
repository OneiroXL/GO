package Master

import (
	"Builder/Director/C/Entity"
	"Builder/Director/I"
)

/*
指挥：核心 负责构建一个工程 如何构建 由他指挥
*/
type Master struct{

}


/*
指挥制造一个产品
*/
func (this *Master) Build(builder I.IBuilder) *Entity.Product {
	builder.BuilderA()
	builder.BuilderB()
	builder.BuilderC()
	builder.BuilderD()

	return builder.GetProduct()
}
package Computer

import (
	"Bridge/I"
)

type Computer struct{
	Brand I.IBrand
}

func (this *Computer) CInfo() {
	this.Brand.Info()
}
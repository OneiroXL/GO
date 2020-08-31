package F

import (
	"Factory/Methon/I"
	"Factory/Methon/C"
)

type TeslaFactory struct{

}

/*
方法一
*/
func (this *TeslaFactory) GetCar() I.ICar {
	var car I.ICar = new(C.Tesla)
	return car
}

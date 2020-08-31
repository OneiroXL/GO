package F

import (
	"Factory/Methon/I"
	"Factory/Methon/C"
)

type WuLingFactory struct{

}

/*
方法一
*/
func (this *WuLingFactory) GetCar() I.ICar {
	var car I.ICar = new(C.WuLin)
	return car
}

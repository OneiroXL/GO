package F

import (
	"Factory/Simple/I"
	"Factory/Simple/C"
)

/*
工厂模式(创造对象，修饰对象)
*/


/*
方法一

不符合开放关闭原则
*/
func GetCar(car string) I.Car {
	switch car {
		case "W":{
			var car I.Car = new(C.WuLin)
			return car
		}
		case "T":{
			var car I.Car = new(C.Tesla)
			return car
		}
	}
	return nil;
}
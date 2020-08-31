package C

import (
	_"Factory/Methon/I"
	"fmt"
)

/*
五菱
*/
type WuLin struct {

}

func (this *WuLin) GetName(){
	fmt.Println("I am WuLin");
}

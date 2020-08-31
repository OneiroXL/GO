package C

import (
	_"Factory/Simple/I"
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

package C

import (
	_"Factory/Methon/I"
	"fmt"
)

/*
特斯拉
*/
type Tesla struct{

}

func (this *Tesla) GetName(){
	fmt.Println("I am Tesla");
}

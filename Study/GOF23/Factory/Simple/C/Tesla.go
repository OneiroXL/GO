package C

import (
	_"Factory/Simple/I"
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

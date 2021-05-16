package Brand

import (
	_"Bridge/I"
	"fmt"
)

type Apple struct {

}

func (this *Apple)Info()  {
	fmt.Print("苹果品牌")
}
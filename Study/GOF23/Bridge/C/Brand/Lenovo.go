package Brand

import (
	_"Bridge/I"
	"fmt"
)

type Lenovo struct {

}

func (this *Lenovo) Info() {
	fmt.Print("联想品牌")
}
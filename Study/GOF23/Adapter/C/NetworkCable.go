package C

import (
	"fmt"
)

/*
要被适配的类
*/
type NetworkCable struct {

}

/*
请求
*/
func (this *NetworkCable) Request(){
	fmt.Println("上网成功")
}

package C

import (
	_"Adapter/I"
)


/*
真真的适配器

---------对象适配器--------

*/
type Adapter struct {
	networkCable *NetworkCable
}

func NewAdapter() *Adapter {
	return &Adapter{
		networkCable:new(NetworkCable),
	}
}

/*
处理请求
*/
func (this *Adapter) HandleRequest() {
	this.networkCable.Request()
}
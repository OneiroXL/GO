package C

import (
	"Adapter/I"
)

type Computer struct {
	
}

/*
上网
*/
func (this *Computer) Net(netToUsb I.INetToUsb){
	netToUsb.HandleRequest()
}
package main

import (
	"Adapter/C"
)

func main(){
	computer := new(C.Computer)
	//networkCable := new(C.NetworkCable)
	adapter := C.NewAdapter()

	computer.Net(adapter)
}



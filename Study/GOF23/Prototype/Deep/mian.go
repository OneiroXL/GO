
package main

import (
	"Prototype/Deep/C"
	"bytes"
	"fmt"
	"encoding/gob"
)


func main(){
	house1 := C.NewHouse()

	house1.Name = "zxl House"
	house1.Size = 127
	house1.Money = 100

	house2 := C.NewHouse()
	deepCopy(house2,house1)


	house2.Size = 300

	fmt.Printf("house1 的地址%p \nhouse1 的内容:%v \n",house1,house1)
	fmt.Printf("house2 的地址%p \nhouse2 的内容:%v \n",house2,house2)
}


func deepCopy(dst, src interface{}) error {
    var buf bytes.Buffer
    if err := gob.NewEncoder(&buf).Encode(src); err != nil {
        return err
    }
    return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
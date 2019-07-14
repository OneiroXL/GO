package tools

import (
	"net"
	"encoding/binary"
	"fmt"
)

type TransferTool struct{
	//链接
	Conn net.Conn
	//缓冲区
	Buffer [1024 * 4]byte
}

func (this *TransferTool) Write(data []byte) error {
	//发送数据长度
	var dataLen uint32 = uint32(len(data) + 8)
	binary.LittleEndian.PutUint32(this.Buffer[0:4],dataLen)
	binary.LittleEndian.PutUint32(this.Buffer[4:8],dataLen)
	binary.LittleEndian.PutUint32(this.Buffer[8:12],uint32(689))
	
	fmt.Println(this.Buffer[0:12])

	_,sendLenErr := this.Conn.Write(this.Buffer[0:12])
	if sendLenErr != nil {
		return sendLenErr
	}
	//发送数据
	_,sendDataErr := this.Conn.Write(data)
	if sendDataErr != nil {
		return sendDataErr
	}
	return nil;
}

func (this *TransferTool) Read() (string,error) {
	dataLen,err := this.Conn.Read(this.Buffer[:])
	fmt.Println(dataLen)
	if(err != nil && dataLen <= 0){
		return "",err
	}
	return string(this.Buffer[0:dataLen]),err;
}
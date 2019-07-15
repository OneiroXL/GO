package tools

import (
	"net"
	"encoding/binary"
	"bytes"
)

type TransferTool struct{
	//链接
	Conn net.Conn
	//缓冲区
	Buffer [1024 * 4]byte
}

func (this *TransferTool) Write(data []byte) error {
	//组装头部
	var dataLen uint32 = uint32(len(data) + 8)
	binary.LittleEndian.PutUint32(this.Buffer[0:4],dataLen)//数据长度
	binary.LittleEndian.PutUint32(this.Buffer[4:8],dataLen)//数据长度
	binary.LittleEndian.PutUint16(this.Buffer[8:10],uint16(689))//请求类型
	//将头部和数据拼接
	var sendData bytes.Buffer
	sendData.Write(this.Buffer[0:12])//这里头部多取两位是因为一个是加密字段，一个是保留字段
	sendData.Write(data)//写入数据
	//发送数据
	_,sendDataErr := this.Conn.Write(sendData.Bytes())
	if sendDataErr != nil {
		return sendDataErr
	}
	return nil;
}

func (this *TransferTool) Read() (string,error) {
	dataLen,err := this.Conn.Read(this.Buffer[:])
	if(err != nil && dataLen <= 0){
		return "",err
	}
	return string(this.Buffer[0:dataLen]),err;
}
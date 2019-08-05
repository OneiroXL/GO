package tools

import (
	"net"
	"encoding/binary"
	"bytes"
)

type TcpTool struct {
	//连接
	Conn net.Conn 
	//缓冲区
	Buffer [1024 * 4]byte
}

/*
写入数据
*/
func (this *TcpTool) Write(data []byte) error {
	//组装头部
	var dateLength uint32 = uint32(len(data) + 4)
	binary.LittleEndian.PutUint32(this.Buffer[0:4],dateLength)//数据长度
	//将头部和数据拼接
	var sendData bytes.Buffer
	sendData.Write(this.Buffer[0:4])//这里头部多取两位是因为一个是加密字段，一个是保留字段
	sendData.Write(data)//写入数据
	//发送数据
	_,sendDataErr := this.Conn.Write(sendData.Bytes())
	if sendDataErr != nil {
		return sendDataErr
	}
	return nil;
}

/*
读取数据
*/
func (this *TcpTool) Read() (string,error) {
	dateLength,err := this.Conn.Read(this.Buffer[:])
	if(err != nil){
		return "",err
	}
	return string(this.Buffer[0:dateLength]),nil
}
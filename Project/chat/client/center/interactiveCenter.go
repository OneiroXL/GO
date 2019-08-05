package center

import (
	"chat/model/base"
	"net"
	"chat/common/tools"
)

/*
交互中心
*/
type InteractiveCenter struct{
	Conn net.Conn
	TcpTool tools.TcpTool
}

/*
处理处理
*/
func (this *InteractiveCenter) InteractiveHandle(interactive base.Interactive) error {
	switch interactive.Type {
		case 100:{
			
		}
	}
	return nil
}


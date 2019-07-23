package process

import (
	"net"
	"fmt"
	"strconv"
	"time"
	"regexp"
	"strings"
	"DouYuClient/common/tools"
)


type BarrageProcess struct{
	conn net.Conn
}

/*
私有变量
*/
var (
	transferTool tools.TransferTool
)

func NewBarrageProcess(conn net.Conn) *BarrageProcess {
	transferTool = tools.TransferTool{
		Conn:conn,
	}
	return &BarrageProcess{
		conn:conn,
	}
}

func (this *BarrageProcess) Start(roomId string){
	this.LoginBarrageServer(roomId)
	this.JoinGroup(roomId)
	go this.KeepLive()
	this.ShowBarrage()
}


/*
登录斗鱼弹幕服务
*/
func (this *BarrageProcess) LoginBarrageServer(roomId string) error {
	var login string = "type@=loginreq/roomid@=" + roomId + "/\x00"
	loginError := transferTool.Write([]byte(login))
	return loginError
}


/*
加入斗鱼弹幕分组
*/
func (this *BarrageProcess) JoinGroup(roomId string) error {
	var joingroup string = "type@=joingroup/rid@=" + roomId + "/gid@=-9999/\x00"
	joingroupError := transferTool.Write([]byte(joingroup))
	return joingroupError
}

/*
展示弹幕
*/
func (this *BarrageProcess) ShowBarrage(){
	for{
		data,err := transferTool.Read();
		if(err != nil){
			fmt.Println("接受数据出错:",err)
			break
		}
		if(this.IsMessageBarrage(data)){
			fmt.Println(this.formatBarrage(data[12:]))
		}
	}
}


/*
格式化弹幕
*/
func (this *BarrageProcess) formatBarrage(barrage string) string {
	userNameReg := regexp.MustCompile(`nn@=(.*)/txt@`)
	userNameByte :=  userNameReg.Find([]byte(barrage))
	messageReg := regexp.MustCompile(`txt@=(.*)/cid@`)
	messageByte := messageReg.Find([]byte(barrage))
	
	userName := strings.Split(string(userNameByte),"@=")
	message := strings.Split(string(messageByte),"@=")
	return userName[1][0:len(userName[1]) - 5] + ":" + message[1][0:len(message[1]) - 5]
}

/*
判断是否消息弹幕
*/
func (this *BarrageProcess) IsMessageBarrage(barrage string) bool {  
	typeReg := regexp.MustCompile(`type@=(.*)/rid@`)
	typeByte := typeReg.Find([]byte(barrage));
	resType := strings.Split(string(typeByte),"@=")

	if(len(typeByte) > 0 && resType[1][0:len(resType[1])- 5] == "chatmsg"){
		return true;
	} else {
		return false
	}
}


/*
心跳机制
*/
func (this *BarrageProcess) KeepLive(){
	for {
		var timestamp string = strconv.Itoa(int(time.Now().Unix()))
		var keepliveInfo string = "type@=keeplive/tick@=" + timestamp + "/\x00"
		err :=  transferTool.Write([]byte(keepliveInfo))
		fmt.Println("一次心跳！")
		if(err != nil){
			fmt.Println("心跳检测阵亡！！！")
			break
		}
		time.Sleep(time.Second * 40)
	}
}

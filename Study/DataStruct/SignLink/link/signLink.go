package link

import(
	"fmt"
)

type SignLink struct {
	Next *SignLink
	Data interface{}
}

/*
添加
*/
func (this *SignLink) Add(data interface{}) {
	temp := this
	for {
		if(temp.Next == nil){
			signLink := &SignLink{
				Data:data,
			}
			temp.Next = signLink 
			return
		}
		temp = temp.Next
	}
}

/*
更新
*/

func (this *SignLink) Update(){

}

/*
移除链表
*/
func (this *SignLink) Remove(data interface{}){
	temp := this
	for {
		if(temp.Next.Data == data){
			temp.Next = temp.Next.Next
			return
		}
		temp = temp.Next
	}
}

/*
打印链表
*/
func (this *SignLink) Show(){
	temp := this
	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}
}

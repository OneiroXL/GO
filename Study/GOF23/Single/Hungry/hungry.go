package Hungry

type Hungry struct{

}

//单利模式-饿汗模式
/*
全局变量上来就创建
可能会浪费内存空间
*/
var (
	hungry Hungry = Hungry{}
)

/*
获取实例
*/
func GetInstance() *Hungry {
	return &hungry
}



package main

import (
	"sync"
)

type Lazy struct{
	

}

/*
单例模式-懒汉模式
*/
var (
	lazy *Lazy
	//look是一个互斥锁
	//Mutex 是互斥的意思
	lock sync.Mutex
)

/*
获取实例
双层检测懒汉式 DCL
*/
func GetInstance() *Lazy {
	if(lazy == nil){
		lock.Lock()
		lazy = new(Lazy)//不是一个原子操作
		/*
		1.分配内存空间
		2.执行构造方法
		3.把这个对象指向这个空间
		*/
		lock.Unlock()
	}
	return lazy
}


func main(){
	var i int = 0
	for(i < 10){
		i++
		go GetInstance()
	}
}

/*
1.反射会破坏单例模式
2.反射不会破坏枚举的单例
*/
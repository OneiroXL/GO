package main

import (
	"fmt"
)

type Node struct {
	X int
	Y int
	Value int
}

func main()  {
	//先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1;
	chessMap[2][3] = 2;
	for _,v := range chessMap{
		for _,v2 := range v{
			fmt.Printf("%v \t",v2)
		}
		fmt.Println()
	}

	var sparseArraySplice []Node
	initNode := Node{}
	initNode.X = 11;
	initNode.Y = 11;
	initNode.Value = 0;
	sparseArraySplice = append(sparseArraySplice,initNode);
	for i,v := range chessMap{
		for j,v2 := range v{
			if v2 != 0 {
				node := Node{}
				node.X = i;
				node.Y = j;
				node.Value = v2;
				sparseArraySplice = append(sparseArraySplice,node);
			}
		}
	}
	

	//输出稀疏数组
	for i,vNode := range sparseArraySplice{
		fmt.Printf("%d: %d %d %d \n",i,vNode.X,vNode.Y,vNode.Value)
	}

	//恢复数组
	var rescoverChessMap [11][11]int
	for _,vNode := range sparseArraySplice[1:len(sparseArraySplice)]{
		rescoverChessMap[vNode.X][vNode.Y] = vNode.Value;
	}

	//看看恢复的数组
	for _,v := range rescoverChessMap{
		for _,v2 := range v{
			fmt.Printf("%v \t",v2)
		}
		fmt.Println()
	}

}



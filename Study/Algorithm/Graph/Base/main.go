package main

import(
	_"./denseGraph"
	"./sparseGraph"
	"math/rand"
	"./model"
	"time"
	"fmt"
)

func main()  {
 	var n int  = 7
	// var m int  = 5
	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包

	// newSparseGraph := sparseGraph.NewSparseGraph(n,false)

	// newSparseGraph.AddEdge(0,1)
	// newSparseGraph.AddEdge(0,2)
	// newSparseGraph.AddEdge(0,5)
	// newSparseGraph.AddEdge(0,6)
	// newSparseGraph.AddEdge(3,4)
	// newSparseGraph.AddEdge(3,5)
	// newSparseGraph.AddEdge(4,5)
	// newSparseGraph.AddEdge(4,6)

	// for k,v := range newSparseGraph.Graph {
	// 	fmt.Printf("和[%d]相连的顶点:",k)
	// 	for l := v.Front(); l != nil; l = l.Next(){
	// 		fmt.Printf("%d,",l.Value)
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println("连通分量为:",newSparseGraph.GetComponentCount())


	sparseWeightGraph := sparseGraph.NewSparseWeightGraph(n,false)


	sparseWeightGraph.AddEdge(0,1,10)
	sparseWeightGraph.AddEdge(0,2,20)
	sparseWeightGraph.AddEdge(0,5,30)
	sparseWeightGraph.AddEdge(0,6,40)
	sparseWeightGraph.AddEdge(3,4,50)
	sparseWeightGraph.AddEdge(3,5,60)
	sparseWeightGraph.AddEdge(4,5,70)
	sparseWeightGraph.AddEdge(4,6,80)

	for k,v := range sparseWeightGraph.Graph {
		fmt.Printf("和[%d]相连的顶点:",k)
		for l := v.Front(); l != nil; l = l.Next(){
			edge := l.Value.(*model.Edge)
			fmt.Printf("%d,权值为:%v;",edge.W(),edge.GetWeight())
		}
		fmt.Println()
	}

	fmt.Println("连通分量为:",sparseWeightGraph.GetComponentCount())

}
package main

import (
	"../Base/sparseGraph"
	"./center"
	_"./help"
	"fmt"
	_"../Model"
)

func main()  {
	sparseWeightGraph := sparseGraph.NewSparseWeightGraph(8,false)


	sparseWeightGraph.AddEdge(4,5,35)
	sparseWeightGraph.AddEdge(4,7,37)
	sparseWeightGraph.AddEdge(5,7,28)
	sparseWeightGraph.AddEdge(0,7,16)
	sparseWeightGraph.AddEdge(1,5,32)
	sparseWeightGraph.AddEdge(0,4,38)
	sparseWeightGraph.AddEdge(2,3,17)
	sparseWeightGraph.AddEdge(1,7,19)

	
	sparseWeightGraph.AddEdge(0,2,26)
	sparseWeightGraph.AddEdge(1,2,36)
	sparseWeightGraph.AddEdge(1,3,29)
	sparseWeightGraph.AddEdge(2,7,34)
	sparseWeightGraph.AddEdge(6,2,40)
	sparseWeightGraph.AddEdge(3,6,52)
	sparseWeightGraph.AddEdge(6,0,58)
	sparseWeightGraph.AddEdge(6,4,93)

	lazyPrim := center.NewLazyPrim(sparseWeightGraph)
	fmt.Println("连通分量为:",sparseWeightGraph.GetComponentCount())
	fmt.Println(lazyPrim.GetMinWeight())

	// heap := help.NewMinHeap(5)

	// edge1 := model.NewEdge(1,2,10)
	// edge2 := model.NewEdge(1,2,30)
	// edge3 := model.NewEdge(1,2,20)
	// edge4 := model.NewEdge(1,2,50)
	// edge5 := model.NewEdge(1,2,40)

	// heap.Push(edge1);
	// heap.Push(edge2);
	// heap.Push(edge3);
	// heap.Push(edge4);
	// heap.Push(edge5);

	// for i := 0; i < 5; i++ {
	// 	v,_ := heap.Pop()
	// 	fmt.Println(v.GetWeight())
	// }

}
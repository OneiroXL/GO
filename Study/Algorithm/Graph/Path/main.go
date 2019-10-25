package main

import (
	"../Base/sparseGraph"
	"./center"
)

func main()  {

	newSparseGraph := sparseGraph.NewSparseGraph(7,false)

	newSparseGraph.AddEdge(0,1)
	newSparseGraph.AddEdge(0,2)
	newSparseGraph.AddEdge(0,5)
	newSparseGraph.AddEdge(0,6)
	newSparseGraph.AddEdge(3,4)
	newSparseGraph.AddEdge(3,5)
	newSparseGraph.AddEdge(4,5)
	newSparseGraph.AddEdge(4,6)

	// path,_ := center.NewPath(newSparseGraph,0)
	
	// path.ShowPath(6)

	shortestPath,_ := center.NewShortestPath(newSparseGraph,0)
	shortestPath.ShowPath(4)
	
}
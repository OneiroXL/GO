package sparseGraph

import (
	"container/list"
	"errors"
	"../../Model"
)

type SparseWeightGraph struct {
	N int
	M int
	Directed bool
	Graph map[int]*list.List
}

/*
初始化稀疏图
*/
func NewSparseWeightGraph(n int,directed bool ) *SparseWeightGraph {
	sparseWeightGraph := &SparseWeightGraph{
		M:0,
		N:n,
		Directed:directed,
		Graph: make(map[int]*list.List),
	}
	for i := 0;i < n; i++{
		sparseWeightGraph.Graph[i] = list.New()
	}
	return sparseWeightGraph
}

/*
边
*/
func (this *SparseWeightGraph) Edge() int {
	return this.M
}
/*
顶点
*/
func (this *SparseWeightGraph) Vertex() int {
	return this.N
}

/*
添加边
*/
func (this *SparseWeightGraph) AddEdge(v int, w int,weight interface{}) (bool,error) {
	if (v < 0 || v > this.N || w < 0 || w > this.N){
		return false,errors.New("已越界")
	}
	isHaveEdge,_ := this.HaveEdge(v,w)
	if(isHaveEdge){
		return true,nil
	}
	edge := model.NewEdge(v,w,weight)
	this.Graph[v].PushBack(edge)

	if(v != w && !this.Directed){
		edge := model.NewEdge(w,v,weight)
		this.Graph[w].PushBack(edge)
	}
	this.M++
	return true,nil
}

/*
是否有边
*/
func (this *SparseWeightGraph) HaveEdge(v int,w int) (bool,error){
	if (v < 0 || v > this.N || w < 0 || w > this.N){
		return false,errors.New("已越界")
	}
	for l := this.Graph[v].Front(); l != nil; l = l.Next(){
		if edge := l.Value.(*model.Edge); edge.W() == w {
			return true,nil
		}
	}
	return false,nil
}


/*
获取连通分量数
*/
func (this *SparseWeightGraph) GetComponentCount() int {
	var visited map[int]bool = make(map[int]bool, this.N)
	var count int

	for i := 0; i < this.N; i++ {
		visited[i] = false
	}

	for i := 0; i < this.Vertex(); i++ {
		if(!visited[i]){
			this.DepthFirst(i,visited)
			count++
		}
	}

	return count
}

/*
深度优先遍历
*/
func (this *SparseWeightGraph) DepthFirst(v int,visited map[int]bool)  {
	visited[v] = true
	elements := this.Graph[v]
	for l := elements.Front(); l != nil; l = l.Next() {
		element := l.Value.(*model.Edge)
		if(!visited[element.W()]){
			this.DepthFirst(element.W(),visited)
		}
	}
}

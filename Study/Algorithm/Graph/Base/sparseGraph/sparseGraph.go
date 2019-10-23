package sparseGraph

import (
	"container/list"
	"errors"
)

type SparseGraph struct {
	N int
	M int
	Directed bool
	Graph map[int]*list.List
}

/*
初始化稀疏图
*/
func NewSparseGraph(n int,directed bool ) *SparseGraph {
	sparseGraph := &SparseGraph{
		M:0,
		N:n,
		Directed:directed,
		Graph: make(map[int]*list.List),
	}
	for i := 0;i < n; i++{
		sparseGraph.Graph[i] = list.New()
	}
	return sparseGraph
}

/*
边
*/
func (this *SparseGraph) Edge() int {
	return this.M
}
/*
顶点
*/
func (this *SparseGraph) Vertex() int {
	return this.N
}

/*
添加边
*/
func (this *SparseGraph) AddEdge(v int, w int) (bool,error) {
	if (v < 0 || v > this.N || w < 0 || w > this.N){
		return false,errors.New("已越界")
	}
	isHaveEdge,_ := this.HaveEdge(v,w)
	if(isHaveEdge){
		return true,nil
	}
	this.Graph[v].PushBack(w)
	if(v != w && !this.Directed){
		this.Graph[w].PushBack(v)
	}
	this.M++
	return true,nil
}

/*
是否有边
*/
func (this *SparseGraph) HaveEdge(v int,w int) (bool,error){
	if (v < 0 || v > this.N || w < 0 || w > this.N){
		return false,errors.New("已越界")
	}
	for l := this.Graph[v].Front(); l != nil; l = l.Next(){
		if(l.Value == w){
			return true,nil
		}
	}
	return false,nil
}


/*
获取连通分量数
*/
func (this *SparseGraph) GetComponentCount() int {
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
func (this *SparseGraph) DepthFirst(v int,visited map[int]bool)  {
	visited[v] = true
	elements := this.Graph[v]
	for l := elements.Front(); l != nil; l = l.Next() {
		element := l.Value.(int)
		if(!visited[element]){
			this.DepthFirst(element,visited)
		}
	}
}

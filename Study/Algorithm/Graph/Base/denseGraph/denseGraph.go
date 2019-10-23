package denseGraph


import (
	"errors"
)

type DenseGraph struct{
	N int
	M int
	Directed bool
	Graph [4][4]bool
}

/*
初始化稠密图
*/

func NewDenseGraph(n int,directed bool) *DenseGraph {
	denseGraph := &DenseGraph{
		N:n,
		M:0,
		Directed:directed,
	}
	for i := 0;i < n; i++ {
		for j := 0;j < n; j++{
			denseGraph.Graph[i][j] = false
		}
	}
	return denseGraph
}

/*
返回多少顶点
*/
func (this *DenseGraph) Vertex() int {
	return this.N
}

/*
返回多少边
*/
func (this *DenseGraph) Edge() int {
	return this.M
}

/*
添加边
*/
func (this *DenseGraph) AddEdge(v int, w int) (bool,error) {
	if (v < 0 || v > this.N || w < 0 || w > this.N){
		return false,errors.New("已越界")
	}
	//检查是否有边
	isHaveEdge,_ := this.HaveEdge(v,w)
	if(isHaveEdge){
		return true,nil
	}
	//添加边
	this.Graph[v][w] = true;
	//若是无向图
	if(!this.Directed){
		this.Graph[w][v] = true;
	}
	this.M++
	return true,nil
}

/*
是否有边
*/
func (this *DenseGraph) HaveEdge(v int, w int) (bool,error) {
	if (v < 0 || v > this.N || w < 0 || w > this.N){
		return false,errors.New("已越界")
	}
	return this.Graph[v][w],nil
}

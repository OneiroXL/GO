package denseGraph


import (
	"errors"
	"../../Model"
)

type DenseWeightGraph struct{
	N int
	M int
	Directed bool
	Graph [4][4]*model.Edge
}

/*
初始化稠密图
*/

func NewDenseWeightGraph(n int,directed bool) *DenseWeightGraph {
	denseWeightGraph := &DenseWeightGraph{
		N:n,
		M:0,
		Directed:directed,
	}
	for i := 0;i < n; i++ {
		for j := 0;j < n; j++{
			denseWeightGraph.Graph[i][j] = nil
		}
	}
	return denseWeightGraph
}

/*
返回多少顶点
*/
func (this *DenseWeightGraph) Vertex() int {
	return this.N
}

/*
返回多少边
*/
func (this *DenseWeightGraph) Edge() int {
	return this.M
}

/*
添加边
*/
func (this *DenseWeightGraph) AddEdge(v int, w int,weight interface{}) (bool,error) {
	if (v < 0 || v > this.N || w < 0 || w > this.N){
		return false,errors.New("已越界")
	}
	//检查是否有边
	isHaveEdge,_ := this.HaveEdge(v,w)
	if(isHaveEdge){
		return true,nil
	}
	//添加边
	edge := model.NewEdge(v,w,weight)
	this.Graph[v][w] = edge;
	//若是无向图
	if(!this.Directed){
		edge := model.NewEdge(v,w,weight)
		this.Graph[w][v] = edge;
	}
	this.M++
	return true,nil
}

/*
是否有边
*/
func (this *DenseWeightGraph) HaveEdge(v int, w int) (bool,error) {
	if (v < 0 || v > this.N || w < 0 || w > this.N){
		return false,errors.New("已越界")
	}
	return this.Graph[v][w] != nil,nil
}

package center

import (
	"../../Base/sparseGraph"
	"../help"
	"container/list"
	"../../Model"
)

type LazyPrim struct {
	graph *sparseGraph.SparseWeightGraph
	pq *help.MinHeap
	marked [100]bool
	mst *list.List
	weight int
}


/*
新建一个切分
*/
func NewLazyPrim(g *sparseGraph.SparseWeightGraph) *LazyPrim {
	lazyPrim := &LazyPrim{
		graph:g,
		pq:help.NewMinHeap(g.Edge()),
		mst:list.New(),
		weight:0,
	}
	for index := 0; index < g.Vertex(); index++ {
		lazyPrim.marked[index] = false
	}
	lazyPrim.StartLazyPrim(0)

	for l := lazyPrim.mst.Front(); l != nil; l = l.Next() {
		lazyPrim.weight += l.Value.(*model.Edge).GetWeight().(int)
	}

	return lazyPrim
}

/*
获取最小生成树
*/
func (this *LazyPrim) GetMst() *list.List {
	return this.mst;
}

/*
获取最小权值
*/
func (this *LazyPrim) GetMinWeight() int {
	return this.weight
}

/*
LazyPrim
*/
func (this *LazyPrim) StartLazyPrim(v int) {
	this.visit(v)
	for (!this.pq.IsEmpty()) {
		edge,_ := this.pq.Pop()
		if(this.marked[edge.V()] == this.marked[edge.W()]){
			continue
		}
		this.mst.PushBack(edge)
		if(!this.marked[edge.V()]){
			this.visit(edge.V());
		} else {
			this.visit(edge.W());
		}
	}
}

/*
visit
*/
func (this *LazyPrim) visit(v int)  {
	if(!this.marked[v]){
		this.marked[v] = true
		for l := this.graph.Graph[v].Front(); l != nil; l = l.Next(){
			edge := l.Value.(*model.Edge)
			if(!this.marked[edge.Other(v)]){
				this.pq.Push(edge)
			}
		}
	}
}
package center

import (
	"errors"
	"fmt"
	"container/list"
	"../../Base/sparseGraph"
)

type ShortestPath struct {
	graph *sparseGraph.SparseGraph
	visited map[int]bool
	s int
	from [7]int
	ord [7]int
}

var ( 
	needChan chan int
)

/*
初始化
*/
func NewShortestPath(graph *sparseGraph.SparseGraph,s int) (*ShortestPath,error) {
	if(s < 0 || s > graph.N ){
		return nil,errors.New("已越界")
	}
	shortestPath := &ShortestPath{
		graph:graph,
		visited:make(map[int]bool, graph.N),
		s:s,
	}
	//初始化值
	for i := 0; i < graph.N; i++ {
		shortestPath.visited[i] = false
		shortestPath.from[i] = -1
		shortestPath.ord[i] = -1
	}
	needChan = make(chan int, graph.N)
	shortestPath.spanFirst(s)

	return shortestPath,nil
}

/*
是否有路径
*/
func (this *ShortestPath) HavePath(w int) (bool,error) {
	if(w < 0 || w > this.graph.N ){
		return false,errors.New("已越界")
	}
	return this.visited[w],nil
}

/*
获取一条路经
*/
func (this *ShortestPath) GetPath(w int) (*list.List,error) {
	if(w < 0 || w > this.graph.N ){
		return nil,errors.New("已越界")
	}
	listPath := list.New()
	p := w
	for p != -1 {
		listPath.PushFront(p)
		p = this.from[p]
	}
	return listPath,nil
}

/*
打印路径
*/
func (this *ShortestPath) ShowPath(w int) error {
	if(w < 0 || w > this.graph.N ){
		return errors.New("已越界")
	}
	listPath,_ := this.GetPath(w)

	for l := listPath.Front();l != nil; l = l.Next(){
		fmt.Printf("%d->",l.Value)
	}
	return nil
}

/*
获取欧最短路劲长度
*/
func(this *ShortestPath) GetLength(w int) (int,error) {
	if(w < 0 || w > this.graph.N ){
		return 0,errors.New("已越界")
	}
	return this.ord[w],nil
}


/*
广度优先遍历
*/
func (this *ShortestPath) spanFirst(s int) {
	needChan <-s
	this.visited[s] = true
	this.ord[s] = 0

	for len(needChan) != 0 {
		var v int = <- needChan
		elements := this.graph.Graph[v]
		for l := elements.Front(); l != nil; l = l.Next() {
			var element int = l.Value.(int)
			if(!this.visited[element]){
				needChan <- element
				this.visited[element] = true
				this.from[element] = v
				this.ord[element] = this.ord[v] + 1
			}
		}
	}
}
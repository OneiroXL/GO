package center

import (
	"errors"
	"fmt"
	"container/list"
	"../../Base/sparseGraph"
)

type Path struct {
	graph *sparseGraph.SparseGraph
	visited map[int]bool
	s int
	from [7]int
}

/*
初始化
*/
func NewPath(graph *sparseGraph.SparseGraph,s int) (*Path,error) {
	if(s < 0 || s > graph.N ){
		return nil,errors.New("已越界")
	}
	path := &Path{
		graph:graph,
		visited:make(map[int]bool, graph.N),
		s:s,
	}
	for i := 0; i < graph.N; i++ {
		path.visited[i] = false
		path.from[i] = -1
	}
	
	//深度遍历
	path.depthFirst(s)

	return path,nil
}


/*
是否有路径
*/
func (this *Path) HavePath(w int) (bool,error) {
	if(w < 0 || w > this.graph.N ){
		return false,errors.New("已越界")
	}
	return this.visited[w],nil
}

/*
获取一条路经
*/
func (this *Path) GetPath(w int) (*list.List,error) {
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
func (this *Path) ShowPath(w int) error {
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
深度优先遍历
*/
func (this *Path) depthFirst(v int) {
	this.visited[v] = true
	elements := this.graph.Graph[v]
	for l := elements.Front(); l != nil; l = l.Next() {
		element := l.Value.(int)
		if(!this.visited[element]){
			this.from[element] = v
			this.depthFirst(element)
		}
	}
}

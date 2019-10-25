package model

type Edge struct {
	A int
	B int
	Weight interface{}
}

/*
初始化
*/
func NewEdge(a int,b int,weight interface{}) *Edge {
	edge := &Edge{
		A:a,
		B:b,
		Weight:weight,
	}
	return edge
}

/*

*/
func (this *Edge) V() int {
	return this.A;
}


/*

*/
func (this *Edge) W() int {
	return this.B;
}

/*
返回另一个边
*/
func (this *Edge) Other(x int) int {
	if(x == this.A){
		return this.B
	}
	return this.A
}

/*
获取权值
*/
func (this *Edge) GetWeight() interface{} {
	return this.Weight
}

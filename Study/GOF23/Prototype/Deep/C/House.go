package C

type House struct{
	Size int
	Money int
	Name string
}



/*
new 一个房间
*/
func NewHouse() * House{
	return &House{}
}

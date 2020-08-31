package Entity

import (
	"fmt"
)

type Product struct{
	A string
	B string
	C string
	D string
}


func (this *Product) SetA(v string)  {
	this.A = v
}

func (this *Product) SetB(v string)  {
	this.B = v
}

func (this *Product) SetC(v string)  {
	this.C = v
}

func (this *Product) SetD(v string)  {
	this.D = v
}

func (this *Product) ToString()  {
	fmt.Println(this)
}
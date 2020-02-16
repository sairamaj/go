package main
import (
	"fmt"
)

type Disc struct{

}
type Toh struct{

}

func (t Toh) Draw(){
	for i:=0 ; i<4; i++{	
		fmt.Println("\t   |")
	}

	d:= Disc{}
	d.Draw()
}

func (t Disc) Draw(){
	fmt.Println("\t------")
}
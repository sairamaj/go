package main
import (
	"fmt"
)

type Disc struct{

}
type Toh struct{
	Discs [3]Disc
}

func (t Toh) Draw(){
	for i:=0 ; i<4; i++{	
		fmt.Println("\t   |")
	}

	d:= Disc{}
	d.Draw()
}

func (t Toh) AddDisc(disc Disc){
	
}

func (t Disc) Draw(){
	fmt.Println("\t------")
}
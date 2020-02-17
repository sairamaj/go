package main

import (
	"fmt"
)

type Disc struct {
	Size  int
}

type Tower struct {
	discs []Disc
}

type Gamer struct {
	towers []Tower	
}

func (t Tower) Draw() {
	for i := 0; i < 4; i++ {
		fmt.Println("\t        |")
	}

	for _, d := range t.discs {
		d.Draw()
	}
}

func (t *Tower) AddDisc(disc Disc) {
	if t.discs == nil {
		t.discs = []Disc{}
	}

	t.discs = append(t.discs, disc)
}

func(t *Tower) RemoveDisc() Disc{
	topDisc := t.discs[len(t.discs)-1]
	t.discs = t.discs[0:len(t.discs)-1]
	return topDisc
}

func (d Disc) Draw() {
	fmt.Print("\t")
	for i :=0 ; i < d.Size; i++{
		fmt.Print("-")
	}
	fmt.Println()
}

func CreateGamer() *Gamer {

	towers := []Tower{Tower{},Tower{},Tower{}}
	towers[0].AddDisc(Disc{Size:12})
	towers[0].AddDisc(Disc{Size:18})
	towers[0].AddDisc(Disc{Size:24})

	g := Gamer{towers:towers}
	return &g
}

func (g *Gamer) Draw(){
	for _, t := range g.towers{
		t.Draw()
	}
}

func(g *Gamer) Move(from int , to int){
	disc := g.towers[from-1].RemoveDisc()
	g.towers[to-1].AddDisc(disc)
	g.Draw()
}
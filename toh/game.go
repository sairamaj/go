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

	for i :=len(t.discs)-1 ; i >=0 ; i-- {
		t.discs[i].Draw()
	}
}

func (t *Tower) addDisc(disc Disc) {
	if t.discs == nil {
		t.discs = []Disc{}
	}else{
		topDisc := t.discs[len(t.discs)-1]
		if topDisc.Size < disc.Size {
			panic ("Disc cannot be added")
		}
	}

	t.discs = append(t.discs, disc)
}

func(t *Tower) removeDisc() Disc{
	if len(t.discs) == 0 {
		panic("Tower does not have discs.")
	}

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
	towers[0].addDisc(Disc{Size:24})
	towers[0].addDisc(Disc{Size:18})
	towers[0].addDisc(Disc{Size:12})

	g := Gamer{towers:towers}
	return &g
}

func (g *Gamer) Draw(){
	for _, t := range g.towers{
		t.Draw()
	}
}

func(g *Gamer) Move(from int , to int) (err error){
	defer func(){
		if r := recover(); r!= nil {
			err = fmt.Errorf("Error (cannot remove) %s (Tower:%d)",  r, from)
		}
	}()

	disc := g.towers[from-1].removeDisc()

	defer func(){
		if r := recover(); r!= nil {
			err = fmt.Errorf("Error(cannot add) %s (Tower:%d)",  r, from)
			g.towers[from-1].addDisc(disc)	// put it back.
		}
	}()

	g.towers[to-1].addDisc(disc)
	g.Draw()

	return err
}
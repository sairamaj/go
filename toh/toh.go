package main

import (
	"fmt"
	"os"
)

func main() {

	ok := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	g := CreateGamer()
	g.Draw()
	for {

		var from int
		var to int
		fmt.Print("Plase enter from and to (ex: 1 2 ):")
		_, err := fmt.Fscanln(os.Stdin, &from, &to)
		ok(err)
		fmt.Println("from:", from, "to:", to)
		g.Move(from, to)
	}
}

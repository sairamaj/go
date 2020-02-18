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
		done, err := g.Move(from, to)
		if err != nil{
			fmt.Println("Error", err)
			g.Draw()
		}
		
		if done {
			fmt.Println("______________________")
			fmt.Println("Game Over!")
			fmt.Println("______________________")
			break
		}
	}
}

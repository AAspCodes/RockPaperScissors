package main

import (
	"fmt"
	"math/rand"
	)



const (
	win = "You Win!!"
	lose = "You Lose!!"
	draw = "It's a Draw..."
)


func main(){
	keepPlaying := true
	for keepPlaying {
		gameLoop()
		
		fmt.Print("Do you want to play again? (y/n): ")
		var answer string
		fmt.Scanf("%s", &answer)

		if answer == "y" {
			keepPlaying = true
		} else {
			keepPlaying = false
		}
	}
	fmt.Println("Bye Bye!!")
}


func gameLoop(){
	var playerMove string
	fmt.Print("Rock, Paper, or Scissors? (r/p/s): ")
	fmt.Scanf("%s", &playerMove)

	var compMove string
	
	switch rand.Intn(3){
	case 0: compMove = "r"
	case 1: compMove = "p"
	case 2: compMove = "s"
	}

	fmt.Println("The Computer Played", compMove)
	
	var output string

	if playerMove == "r"{
		switch compMove{
		case "r": output = draw
		case "p": output = lose
		case "s": output = win
		}
	} else if playerMove == "p"{
		switch compMove{
		case "r": output = win
		case "p": output = draw
		case "s": output = lose
		}
	} else {
		switch compMove{
		case "r": output = lose
		case "p": output = win
		case "s": output = draw
		}
	}
	fmt.Println(output)
}

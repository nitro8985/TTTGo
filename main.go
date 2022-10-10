package main

import "fmt"

func main() {
	game := new(ttt_game)
	for Play(game) {
	}
}

func Play(game *ttt_game) bool {
	game.Init()
	game.printField()
	for !game.finished && !game.tie {
		game.UpdateState()
	}
	fmt.Println("Wanna play again?(y/n)")
	var answer string
	_, err := fmt.Scan(&answer)
	for err != nil {
		fmt.Println("Wrong answer=/")
		_, err = fmt.Scan(&answer)
	}
	if answer == "y" {
		return true
	} else {
		fmt.Println("Thank you! Good bye!")
		return false
	}
}

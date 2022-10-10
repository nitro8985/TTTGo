package main

import (
	"fmt"
	"math/rand"
)

type ttt_game struct {
	field    [3][3]byte
	finished bool
	tie      bool
	curr     bool
	mode     bool
	diff     int
}

func (game *ttt_game) Init() {
	game.finished = false
	game.tie = false
	game.curr = true
	game.mode = false
	fmt.Println("Choose game mode:\n1 - PvP\n2 - PvE")
	var md int
	_, err := fmt.Scan(&md)
	for err != nil && (md == 1 || md == 2) {
		_, err = fmt.Scan(&md)
	}
	if md == 2 {
		game.mode = true
		fmt.Println("Choose difficulty:\n1 - easy\n2 - nightmare")
		var dif int
		_, erro := fmt.Scan(&dif)
		for erro != nil && (dif == 1 || dif == 2) {
			_, erro = fmt.Scan(&dif)
		}
		switch dif {
		case 1:
			game.diff = 1
		case 2:
			game.diff = 5
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.field[i][j] = '*'
		}
	}
	fmt.Println("Randomizing turn...")
	rand := rand.Intn(1000)
	if rand%2 == 0 {
		game.curr = false
	}
}

func (game *ttt_game) printField() {
	fmt.Println("  1 2 3")
	for i := 0; i < 3; i++ {
		fmt.Print(i+1, " ")
		for j := 0; j < 3; j++ {
			fmt.Printf("%c ", game.field[i][j])
		}
		fmt.Println()
	}
}

func (game *ttt_game) checkState() {
	hasTurnsLeft := false
	hasWinner := false
	diag1 := game.field[0][0]
	diag2 := game.field[2][0]
	for i := 0; i < 3; i++ {
		currRow := game.field[i][0]
		currCol := game.field[0][i]
		countRows, countCols, countDiag1, countDiag2 := 0, 0, 0, 0
		for j := 0; j < 3; j++ {
			if game.field[i][j] == '*' {
				hasTurnsLeft = true
			}
			if game.field[i][j] == currRow && game.field[i][j] != '*' {
				countRows++
			}
			if game.field[j][i] == currCol && game.field[j][i] != '*' {
				countCols++
			}
			if diag1 == game.field[j][j] && game.field[j][j] != '*' {
				countDiag1++
			}
			if diag2 == game.field[2-j][j] && game.field[2-j][j] != '*' {
				countDiag2++
			}
			if countRows == 3 || countCols == 3 || countDiag1 == 3 || countDiag2 == 3 {
				hasWinner = true
			}
		}
	}
	if hasWinner {
		var c byte
		if game.curr {
			c = 'O'
		} else {
			c = 'X'
		}
		game.finished = true
		fmt.Printf("%c won!\n", c)
		return
	}
	if !hasTurnsLeft {
		game.tie = true
		fmt.Print("Tie!\n")
	}
}

func (game *ttt_game) UpdateState() {
	var a, b int
	var c byte
	if game.curr {
		c = 'X'
	} else {
		c = 'O'
	}
	if game.curr {
		fmt.Printf("%c turn\nEnter row & col\n", c)
		args, err := fmt.Scan(&a, &b)
		for err != nil && args != 2 {
			fmt.Println("Check input")
			a, b = -1, -1
			_, err = fmt.Scan(&a, &b)
		}
		if a < 4 && b < 4 && a > 0 && b > 0 && game.field[a-1][b-1] == '*' {
			game.field[a-1][b-1] = c
			game.curr = !game.curr
		} else {
			fmt.Println("Wrong input")
		}
		game.checkState()
	} else {
		if game.mode && !game.finished && !game.tie {
			var comp AI
			if game.curr {
				comp.own = 'X'
				comp.opp = 'O'
			} else {
				comp.own = 'O'
				comp.opp = 'X'
			}
			row, col := comp.MakeDecision(game)
			game.field[row][col] = comp.own
			game.curr = !game.curr
		}
		game.checkState()
	}
	game.printField()
}

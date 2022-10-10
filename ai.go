package main

type AI struct {
	own, opp byte
}

func moveRemains(game *ttt_game) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.field[i][j] == '*' {
				return true
			}
		}
	}
	return false
}

func (comp *AI) minimax(game *ttt_game, depth int, isMax bool) int {
	curr_score := comp.evaluateCurrState(game)
	if curr_score == 10 {
		return curr_score
	}
	if curr_score == -10 {
		return curr_score
	}
	if !moveRemains(game) || depth == 0 {
		return 0
	}
	if isMax {
		best_score := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if game.field[i][j] == '*' {
					game.field[i][j] = comp.own
					best_score = max(best_score, comp.minimax(game, depth-1, !isMax))
					game.field[i][j] = '*'
				}
			}
		}
		return best_score
	} else {
		best_score := 1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if game.field[i][j] == '*' {
					game.field[i][j] = comp.opp
					best_score = min(best_score, comp.minimax(game, depth-1, !isMax))
					game.field[i][j] = '*'
				}
			}
		}
		return best_score
	}
}

func (comp *AI) evaluateCurrState(game *ttt_game) int {
	for row := 0; row < 3; row++ {
		counter := 0
		curr := game.field[row][0]
		for i := 0; i < 3; i++ {
			if game.field[row][i] == curr {
				counter++
			}
		}
		if counter == 3 {
			if curr == comp.own {
				return 10
			} else if curr == comp.opp {
				return -10
			}
		}
	}
	for col := 0; col < 3; col++ {
		counter := 0
		curr := game.field[0][col]
		for i := 0; i < 3; i++ {
			if game.field[i][col] == curr {
				counter++
			}
		}
		if counter == 3 {
			if curr == comp.own {
				return 10
			} else if curr == comp.opp {
				return -10
			}
		}
	}
	d1 := game.field[0][0]
	d2 := game.field[2][0]
	c1, c2 := 0, 0
	for i := 0; i < 3; i++ {
		if game.field[i][i] == d1 {
			c1++
		}
	}
	if c1 == 3 {
		if d1 == comp.own {
			return 10
		} else if d1 == comp.opp {
			return -10
		}
	}
	for i := 0; i < 3; i++ {
		if game.field[2-i][i] == d2 {
			c2++
		}
	}
	if c2 == 3 {
		if d2 == comp.own {
			return 10
		} else if d2 == comp.opp {
			return -10
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func (comp *AI) MakeDecision(game *ttt_game) (int, int) {
	depth := game.diff
	best_score := -1000
	bestRow, bestCol := -1, -1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.field[i][j] == '*' {
				game.field[i][j] = comp.own
				curr_move_score := comp.minimax(game, depth, false)
				game.field[i][j] = '*'
				if curr_move_score > best_score {
					bestRow = i
					bestCol = j
					best_score = curr_move_score
				}
			}
		}
	}
	return bestRow, bestCol
}

package app

import "fmt"

type Game struct {
	board [3][3]int
	// need store first and second player?
	firstPlayer   Player
	secondPlayer  Player
	currentPlayer Player
	round         int
	currentMove   Move
}

func (g Game) Finished() bool {
	var endCondition = false

	// check only win condition only for last move made
	for i := 0; i < 3; i++ {
		if g.board[g.currentMove.row][i] != g.currentPlayer.piece {
			break
		}

		if i == 2 {
			// current player wins
			endCondition = true
		}
	}

	for i := 0; i < 3; i++ {
		if g.board[i][g.currentMove.column] != g.currentPlayer.piece {
			break
		}

		if i == 2 {
			// current player wins
			endCondition = true
		}
	}

	if g.currentMove.row == g.currentMove.column {
		// check if we are on a diagonal
		for i := 0; i < 3; i++ {
			if g.board[i][i] != g.currentPlayer.piece {
				break
			}

			if i == 2 {
				// current player wins
				endCondition = true
			}
		}
	}

	if g.currentMove.row+g.currentMove.column == 2 {
		// other diagonal
		for i := 0; i < 3; i++ {
			if g.board[i][2-i] != g.currentPlayer.piece {
				break
			}

			if i == 2 {
				// current player wins
				endCondition = true
			}
		}
	}

	if g.round == 8 {
		// draw
		endCondition = true
	}

	return endCondition
}

func (g Game) Display() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("|")

			switch piece := g.board[i][j]; piece {
			case 1:
				fmt.Print("x")
			case 2:
				fmt.Print("o")
			default:
				fmt.Print(" ")
			}
		}

		fmt.Println("|")
	}
}

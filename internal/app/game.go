package app

type Game struct {
	firstPlayer  Player
	secondPlayer Player
	state        State
}

type State struct {
	board         [3][3]int
	currentPlayer Player
	round         int
	currentMove   Move
}

func (s State) Finished() bool {
	var endCondition = false

	// check only win condition only for last move made
	for i := 0; i < 3; i++ {
		if s.board[s.currentMove.row][i] != s.currentPlayer.piece {
			break
		}

		if i == 2 {
			// current player wins
			endCondition = true
		}
	}

	for i := 0; i < 3; i++ {
		if s.board[i][s.currentMove.column] != s.currentPlayer.piece {
			break
		}

		if i == 2 {
			// current player wins
			endCondition = true
		}
	}

	if s.currentMove.row == s.currentMove.column {
		// check if we are on a diagonal
		for i := 0; i < 3; i++ {
			if s.board[i][i] != s.currentPlayer.piece {
				break
			}

			if i == 2 {
				// current player wins
				endCondition = true
			}
		}
	}

	if s.currentMove.row+s.currentMove.column == 2 {
		// other diagonal
		for i := 0; i < 3; i++ {
			if s.board[i][2-i] != s.currentPlayer.piece {
				break
			}

			if i == 2 {
				// current player wins
				endCondition = true
			}
		}
	}

	if s.round == 8 {
		// draw
		endCondition = true
	}

	return endCondition
}

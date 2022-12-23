package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToIntSlice(ss []string) ([]int, error) {
	si := make([]int, 0, len(ss))

	for _, s := range ss {
		i, err := strconv.Atoi(s)

		if err != nil {
			return si, err
		}

		si = append(si, i)
	}

	return si, nil
}

type move struct {
	row    int
	column int
}

func (m move) validate(board [3][3]int) (bool, error) {
	if m.row > 2 || m.column > 2 || m.row < 0 || m.column < 0 {
		return false, errors.New("Index out of bounds")
	}

	if board[m.row][m.column] != 0 {
		return false, errors.New("Position already filled")
	}

	return true, nil
}

type player struct {
	name  string
	piece int
}

func (p player) askMove() (move, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter move row and column separated by comma: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	moveAsStringArray := strings.Split(text, ",")

	moveAsIntArray, err := stringToIntSlice(moveAsStringArray)

	if err != nil {
		return move{}, err
	}

	move := move{moveAsIntArray[0] - 1, moveAsIntArray[1] - 1}

	return move, nil
}

type game struct {
	board [3][3]int
	// need store first and second player?
	firstPlayer   player
	secondPlayer  player
	currentPlayer player
	round         int
	currentMove   move
}

func (g game) finished() bool {
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

func (g game) display() {
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

func Run() {
	firstPlayer := player{"First", 1}
	secondPlayer := player{"Second", 2}
	game := game{firstPlayer: firstPlayer, secondPlayer: secondPlayer}
	var err error

	for {
		if game.round%2 == 0 {
			game.currentPlayer = game.firstPlayer
		} else {
			game.currentPlayer = game.secondPlayer
		}

		for validMove := false; validMove == false; {
			game.currentMove, err = game.currentPlayer.askMove()

			if err != nil {
				fmt.Println(err)
				return
			}

			validMove, err = game.currentMove.validate(game.board)

			if err != nil {
				fmt.Println(err)
			}
		}

		game.board[game.currentMove.row][game.currentMove.column] = game.currentPlayer.piece

		game.display()

		if game.finished() {
			break
		}

		game.round++
	}
}

package app

import "errors"

type Move struct {
	row    int
	column int
}

func (m Move) validate(board [3][3]int) (bool, error) {
	if m.row > 2 || m.column > 2 || m.row < 0 || m.column < 0 {
		return false, errors.New("Index out of bounds")
	}

	if board[m.row][m.column] != 0 {
		return false, errors.New("Position already filled")
	}

	return true, nil
}

func (m Move) execute(board *[3][3]int, piece int) {
	board[m.row][m.column] = piece
}

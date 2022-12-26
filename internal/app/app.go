package app

import (
	"fmt"
)

func Run() {
	firstPlayer := Player{"First", 1}
	secondPlayer := Player{"Second", 2}
	game := Game{firstPlayer: firstPlayer, secondPlayer: secondPlayer}
	var err error

	for {
		if game.round%2 == 0 {
			game.currentPlayer = game.firstPlayer
		} else {
			game.currentPlayer = game.secondPlayer
		}

		for validMove := false; validMove == false; {
			game.currentMove, err = game.currentPlayer.AskMove()

			if err != nil {
				fmt.Println(err)
				return
			}

			validMove, err = game.currentMove.validate(game.board)

			if err != nil {
				fmt.Println(err)
			}
		}

		game.currentMove.execute(&game.board, game.currentPlayer.piece)

		game.Display()

		if game.Finished() {
			break
		}

		game.round++
	}
}

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
		if game.state.round%2 == 0 {
			game.state.currentPlayer = game.firstPlayer
		} else {
			game.state.currentPlayer = game.secondPlayer
		}

		for validMove := false; validMove == false; {
			game.state.currentMove, err = game.state.currentPlayer.AskMove()

			if err != nil {
				fmt.Println(err)
				return
			}

			validMove, err = game.state.currentMove.validate(game.state.board)

			if err != nil {
				fmt.Println(err)
			}
		}

		game.state.currentMove.execute(&game.state.board, game.state.currentPlayer.piece)

		Display(game.state)

		if game.state.Finished() {
			break
		}

		game.state.round++
	}
}

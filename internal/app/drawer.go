package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Display(s State) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("|")

			switch piece := s.board[i][j]; piece {
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

func AskMove() ([]int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter move row and column separated by comma: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	moveAsStringArray := strings.Split(text, ",")

	moveAsIntArray, err := StringToIntSlice(moveAsStringArray)

	if err != nil {
		return moveAsIntArray, err
	}

	return moveAsIntArray, nil
}

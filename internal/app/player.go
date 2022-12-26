package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	name  string
	piece int
}

func (p Player) AskMove() (Move, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter move row and column separated by comma: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	moveAsStringArray := strings.Split(text, ",")

	moveAsIntArray, err := StringToIntSlice(moveAsStringArray)

	if err != nil {
		return Move{}, err
	}

	move := Move{moveAsIntArray[0] - 1, moveAsIntArray[1] - 1}

	return move, nil
}

package app

type Player struct {
	name  string
	piece int
}

func (p Player) AskMove() (Move, error) {
	moveAsIntArray, err := AskMove()

	if err != nil {
		return Move{}, err
	}

	move := Move{moveAsIntArray[0] - 1, moveAsIntArray[1] - 1}

	return move, nil
}

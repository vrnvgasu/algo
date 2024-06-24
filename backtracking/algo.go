package backtracking

import "fmt"

var counter = 0

func Queens(board Board) bool {
	if board.Has8Queen() {
		fmt.Printf("counter: %d\n", counter)

		return true
	}

	counter++

	board.Draw()
	positions := board.ShufflePositions() // для разных комбинаций
	//positions := board.Positions()
	for _, pos := range positions {
		if ok := board.PlaceQueen(pos); !ok {
			return false
		}

		solution := Queens(board)
		if solution {
			return true
		} else {
			board.RemoveQueen(pos)
		}
	}

	return false
}

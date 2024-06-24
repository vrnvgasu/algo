package backtracking

import (
	"fmt"
	"math/rand"
)

/*
[][][][][][][][]
[][][][][][][][]
[][][][][][][][]
[][][][][][][][]
[][][][][][][][]
[][][][][][][][]
[][][][][][][][]
[][][][][][][][]
*/

type queen struct {
	position
	AttackedPositions []position
}

type position struct {
	X, Y int
}

type Board struct {
	board  [8][8]int
	queues []queen
}

func NewBoard() Board {
	return Board{
		board:  [8][8]int{},
		queues: []queen{},
	}
}

func (b *Board) Has8Queen() bool {
	if len(b.queues) == 0 || len(b.queues) == 8 {
		fmt.Println("SUCCESS")
		b.Draw()
	}

	return len(b.queues) == 8
}

func (b *Board) Positions() []position {
	positions := []position{}
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			if b.board[x][y] == 0 {
				positions = append(positions, position{x, y})
			}
		}
	}

	return positions
}

func (b *Board) ShufflePositions() []position {
	positions := []position{}

	rand.Intn(8)
	xStart := rand.Intn(8)
	yStart := rand.Intn(8)

	for x := xStart; x < 8; x++ {
		for y := yStart; y < 8; y++ {
			if b.board[x][y] == 0 {
				positions = append(positions, position{x, y})
			}
		}
	}
	for x := xStart; x >= 0; x-- {
		for y := yStart; y >= 0; y-- {
			if b.board[x][y] == 0 {
				positions = append(positions, position{x, y})
			}
		}
	}

	return positions
}

func (b *Board) PlaceQueen(p position) bool {
	if b.board[p.X][p.Y] == 1 {
		return false
	}

	q := queen{
		position:          p,
		AttackedPositions: []position{},
	}

	// горизонтально
	for x := 0; x < 8; x++ {
		if b.board[x][p.Y] == 0 {
			q.AttackedPositions = append(q.AttackedPositions, position{
				X: x,
				Y: p.Y,
			})
		}
		b.board[x][p.Y] = 1

	}
	// вертикально
	for y := 0; y < 8; y++ {
		if b.board[p.X][y] == 0 {
			q.AttackedPositions = append(q.AttackedPositions, position{
				X: p.X,
				Y: y,
			})
		}
		b.board[p.X][y] = 1
	}
	// вправо вверх
	y := p.Y
	for x := p.X; x < 8 && y < 8; x++ {
		if b.board[x][y] == 0 {
			q.AttackedPositions = append(q.AttackedPositions, position{X: x, Y: y})
		}
		b.board[x][y] = 1
		y++
	}
	// влево вверх
	y = p.Y
	for x := p.X; x >= 0 && y < 8; x-- {
		if b.board[x][y] == 0 {
			q.AttackedPositions = append(q.AttackedPositions, position{X: x, Y: y})
		}
		b.board[x][y] = 1
		y++
	}
	// вправо вниз
	y = p.Y
	for x := p.X; x < 8 && y >= 0; x++ {
		if b.board[x][y] == 0 {
			q.AttackedPositions = append(q.AttackedPositions, position{X: x, Y: y})
		}
		b.board[x][y] = 1
		y--
	}
	// влево вниз
	y = p.Y
	for x := p.X; x >= 0 && y >= 0; x-- {
		if b.board[x][y] == 0 {
			q.AttackedPositions = append(q.AttackedPositions, position{X: x, Y: y})
		}
		b.board[x][y] = 1
		y--
	}

	b.queues = append(b.queues, q)

	return true
}

func (b *Board) RemoveQueen(p position) {
	queues := []queen{}
	for _, q := range b.queues {
		if p.Y != q.Y && p.X != q.X {
			queues = append(queues, q)
			continue
		}

		for _, ap := range q.AttackedPositions {
			b.board[ap.X][ap.Y] = 0
		}
	}

	b.queues = queues
}

func (b *Board) Draw() {
	fmt.Println("---------------------------")
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			var qu *queen
			for _, q := range b.queues {
				if q.Y == y && q.X == x {
					qu = &q
					break
				}
			}

			if qu == nil {
				if b.board[x][y] == 0 {
					fmt.Print("[ ]")
				} else {
					fmt.Print("[-]")
				}

			} else {
				fmt.Print("[o]")
			}
		}
		fmt.Println()
	}
	fmt.Println("---------------------------")
}

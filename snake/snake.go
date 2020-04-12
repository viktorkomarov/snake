package snake

import (
	"math/rand"
	
	"gitlab.com/VictorKomarov/snake/game"
)

type NextStep string

const (
	Up NextStep = "up"
	Down NextStep = "down"
	Left NextStep = "left"
	Right NextStep = "right"
)

type Snake struct {
	size int
	nextStep NextStep
	head *Node
}

type Node struct {
	coordinate game.Cell
	tail *Node
}

func New(fromX, toX, fromY, toY int)*Snake {
	x := rand.Intn(toX-fromX) + fromX
	y := rand.Intn(toY-fromY) + fromY
	defaultStart := Right
	if x >= fromX - 5 {
		defaultStart = Left
	}
	return &Snake {
		size: 1,
		nextStep : defaultStart,
		head: &Node{
			coordinate : game.Cell{X: x, Y: y},
		},
	}
}

func moveCoordinate(coordinate game.Cell, to NextStep) game.Cell {
	switch to {
	case Up:
		coordinate.X += 1
	case Down:
		coordinate.X -= 1
	case Left:
		coordinate.Y -= 1
	case Right:
		coordinate.Y += 1
	default:
		panic("change logic")		
	}

	return coordinate
}

func (s *Snake) Move() {
	node := s.head
	fromCoordinate := node.coordinate
	node.coordinate = moveCoordinate(node.coordinate, s.nextStep)
	node = node.tail
	for node != nil {
		node.coordinate = fromCoordinate
		node = node.tail
	}
}

func (s *Snake) Snapshot() []game.Cell {
	cells := make([]game.Cell, 0, s.size)
	node := s.head
	for node != nil {
		cells = append(cells, node.coordinate)
		node = node.tail
	}

	return cells
}
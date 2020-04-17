package game

import (
	"math/rand"
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
	coordinate Cell
	tail *Node
}

func NewSnake(fromX, toX, fromY, toY int)*Snake {
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
			coordinate : Cell{X: x, Y: y},
		},
	}
}

func moveCoordinate(coordinate Cell, to NextStep) Cell {
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
		oldCoordinate := node.coordinate
		node.coordinate = fromCoordinate
		node = node.tail
		fromCoordinate = oldCoordinate
	}
}

func (s *Snake) MoveByUser(road NextStep) {
	s.nextStep = road
	s.Move()
}

func (s *Snake) Snapshot() []Cell {
	cells := make([]Cell, 0, s.size)
	node := s.head
	for node != nil {
		cells = append(cells, node.coordinate)
		node = node.tail
	}

	return cells
}

// pointer to last head
func(s *Snake) Eat(food Cell) bool {
	node := s.head
	if !(node.coordinate.X == food.X && node.coordinate.Y == food.Y) {
		return false
	}

	for node.tail != nil {
		node = node.tail
	}
	
	newX, newY := node.coordinate.X, node.coordinate.Y
	switch s.nextStep {
	case Up:
		newX -= 1
	case Down:
		newX += 1	
	case Left:
		newY += 1
	case Right:
		newY -= 1		
	}
	
	node.tail = &Node{
		coordinate :Cell {
			X: newX,
			Y: newY,
		},
		tail: nil,
	}
	s.size++

	return true
}
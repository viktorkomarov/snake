package game

import (
	"math/rand"
)

type NextStep string

const (
	Up    NextStep = "up"
	Down  NextStep = "down"
	Left  NextStep = "left"
	Right NextStep = "right"
)

type Snake struct {
	size     int
	NextStep NextStep
	borders  []int
	head     *Node
}

type Node struct {
	coordinate Cell
	tail       *Node
}

func NewSnake(fromX, toX, fromY, toY int) *Snake {
	x := rand.Intn(toX-fromX) + fromX
	y := rand.Intn(toY-fromY) + fromY
	defaultStart := Right
	if x >= fromX-5 {
		defaultStart = Left
	}
	return &Snake{
		size:     1,
		NextStep: defaultStart,
		head: &Node{
			coordinate: Cell{X: x, Y: y},
		},
		borders: []int{fromX, toX, fromY, toY},
	}
}

// borders [fromX, toX, fromY, toY]
func (s *Snake) moveCoordinate(coordinate Cell, to NextStep) Cell {
	switch to {
	case Up:
		if coordinate.X+1 >= s.borders[1] {
			coordinate.X = s.borders[0]
			return coordinate
		}

		coordinate.X += 1
	case Down:
		if coordinate.X-1 <= s.borders[0] {
			coordinate.X = s.borders[1] - 1
			return coordinate
		}

		coordinate.X -= 1
	case Left:
		if coordinate.Y-1 <= s.borders[2] {
			coordinate.Y = s.borders[3] - 1
			return coordinate
		}

		coordinate.Y -= 1
	case Right:
		if coordinate.Y+1 >= s.borders[3] {
			coordinate.Y = s.borders[2] + 1
			return coordinate
		}

		coordinate.Y += 1
	default:
		panic("change logic")
	}

	return coordinate
}

func (s *Snake) IsEnd() bool {
	headCoordinate := s.head.coordinate
	head := s.head.tail
	for head != nil {
		if s.head.coordinate == headCoordinate {
			return true
		}

		head = s.head.tail
	}

	return false
}

func (s *Snake) move() {
	node := s.head
	fromCoordinate := node.coordinate
	node.coordinate = s.moveCoordinate(node.coordinate, s.NextStep)
	node = node.tail
	for node != nil {
		oldCoordinate := node.coordinate
		node.coordinate = fromCoordinate
		node = node.tail
		fromCoordinate = oldCoordinate
	}
}

func (s *Snake) validateUserRoad(road NextStep) bool {
	if s.size == 1 {
		return true
	}

	switch road {
	case Left:
		return s.NextStep != Right
	case Right:
		return s.NextStep != Left
	case Up:
		return s.NextStep != Down
	case Down:
		return s.NextStep != Up
	}

	return true
}

func (s *Snake) MoveByUser(road NextStep) {
	if s.validateUserRoad(road) {
		s.NextStep = road
	}

	s.move()
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
func (s *Snake) Eat(food Cell) bool {
	node := s.head
	if !(node.coordinate.X == food.X && node.coordinate.Y == food.Y) {
		return false
	}

	for node.tail != nil {
		node = node.tail
	}

	newX, newY := node.coordinate.X, node.coordinate.Y
	switch s.NextStep {
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
		coordinate: Cell{
			X: newX,
			Y: newY,
		},
		tail: nil,
	}
	s.size++

	return true
}

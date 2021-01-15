package game

import "errors"

var ErrSnakeFail = errors.New("snake fail")

type Snake struct {
	score      int
	prevStep   Event
	borders    []int
	head       *Node
	fromX, toX int
	fromY, toY int
}

type Node struct {
	coordinate Cell
	tail       *Node
}

func NewSnake() *Snake {
	return &Snake{}
}

func (s *Snake) Score() int {
	return s.score
}

func (s *Snake) move(step Event, food Cell) error {
	if !validateStep(s.prevStep, step) {
		step = s.prevStep
	}

	prevCoord := s.moveHead(step)
	addNode := s.eat(food)
	current := s.head
	for current.tail != nil {
		current.tail.coordinate, prevCoord = prevCoord, current.tail.coordinate
		if prevCoord == s.head.coordinate {
			return ErrSnakeFail
		}

		current = current.tail
	}

	if addNode {
		current.tail = &Node{coordinate: prevCoord}
	}

	s.prevStep = step
	return nil
}

func (s *Snake) moveHead(step Event) Cell {
	prev := s.head.coordinate

	cell := NewCell(step)
	s.head.coordinate.X += cell.X
	s.head.coordinate.Y += cell.Y

	if s.head.coordinate.X < s.fromX {
		s.head.coordinate.X = s.toX
	}

	if s.head.coordinate.X > s.toX {
		s.head.coordinate.X = s.fromX
	}

	if s.head.coordinate.Y < s.fromY {
		s.head.coordinate.Y = s.toY
	}

	if s.head.coordinate.Y > s.toY {
		s.head.coordinate.Y = s.fromY
	}

	return prev
}

func (s *Snake) eat(food Cell) bool {
	return s.head.coordinate == food
}

func validateStep(prevStep, nextStep Event) bool {
	switch prevStep {
	case Up:
		return nextStep != Down
	case Down:
		return nextStep != Up
	case Left:
		return nextStep != Right
	case Right:
		return nextStep != Left
	}

	return false
}

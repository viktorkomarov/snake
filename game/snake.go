package game

import "errors"

var ErrSnakeFail = errors.New("snake fail")

type Snake struct {
	size     int
	prevStep Event
	Head     *Node
	arena    Arena
}

type Node struct {
	Coordinate Cell
	Tail       *Node
}

func NewSnake(arena Arena, firstStep Event) *Snake {
	return &Snake{
		size:     1,
		prevStep: firstStep,
		arena:    arena,
		Head: &Node{
			Coordinate: arena.RandomCell(),
		},
	}
}

func (s *Snake) Size() int {
	return s.size
}

func (s *Snake) move(step Event, food Cell) (bool, error) {
	if !validateStep(s.prevStep, step) {
		step = s.prevStep
	}

	prevCoord := s.moveHead(step)
	addNode := s.eat(food)
	current := s.Head

	for current.Tail != nil {
		current.Tail.Coordinate, prevCoord = prevCoord, current.Tail.Coordinate
		if prevCoord == s.Head.Coordinate {
			return false, ErrSnakeFail
		}

		current = current.Tail
	}

	if addNode {
		s.size++
		current.Tail = &Node{Coordinate: prevCoord}
	}

	s.prevStep = step
	return addNode, nil
}

func (s *Snake) moveHead(step Event) Cell {
	prev := s.Head.Coordinate
	cell := NewCell(step)
	s.Head.Coordinate.X += cell.X
	s.Head.Coordinate.Y += cell.Y

	if s.Head.Coordinate.X < s.arena.FromX {
		s.Head.Coordinate.X = s.arena.ToX
	}

	if s.Head.Coordinate.X > s.arena.ToX {
		s.Head.Coordinate.X = s.arena.FromX
	}

	if s.Head.Coordinate.Y < s.arena.FromY {
		s.Head.Coordinate.Y = s.arena.ToY
	}

	if s.Head.Coordinate.Y > s.arena.ToY {
		s.Head.Coordinate.Y = s.arena.FromY
	}

	return prev
}

func (s *Snake) eat(food Cell) bool {
	return s.Head.Coordinate == food
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

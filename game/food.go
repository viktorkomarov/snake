package game

import (
	"math/rand"
)

type Food struct {
	actual     Cell
	fromX, toX int
	fromY, toY int
}

func NewFood(fromX, toX, fromY, toY int) Food {
	f := Food{
		fromX: fromX,
		toX:   toX,
		fromY: fromY,
		toY:   toY,
	}
	f.regenerate()
	return f
}

func (f *Food) regenerate() {
	f.actual.X = rand.Intn(f.toX-f.fromX) + f.fromX
	f.actual.Y = rand.Intn(f.toY-f.fromY) + f.fromY
}

func (f *Food) Coordinate() Cell {
	return f.actual
}

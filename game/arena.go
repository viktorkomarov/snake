package game

type Cell struct {
	X, Y int
}

func NewCell(step Event) Cell {
	switch step {
	case Up:
		return Cell{0, 1}
	case Down:
		return Cell{0, -1}
	case Left:
		return Cell{-1, 0}
	case Right:
		return Cell{0, 1}
	}

	return Cell{0, 0}
}

type Arena struct {
	FromX, ToX int
	FromY, ToY int
}

func NewArena(width, height int) Arena {
	x, y := width/2, height/2
	return Arena{
		FromX: x - (x / 2),
		ToX:   x + (x / 2),
		FromY: y - (y / 2),
		ToY:   y + (y / 2),
	}
}

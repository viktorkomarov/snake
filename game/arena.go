package game

import (
	"github.com/nsf/termbox-go"
)

type Config struct {
	snakeColor termbox.Attribute
	foodColor termbox.Attribute
	bgColor termbox.Attribute
}

type Cell struct {
	X, Y int
}

func defaultConfig() Config {
	return Config {
		snakeColor  : termbox.ColorGreen,
		foodColor   : termbox.ColorRed,
		bgColor     : termbox.ColorBlue,
	}
}

type Arena struct {
	FromX, ToX int
	FromY, ToY int
	cfg Config
}

func NewArena(cfg *Config) *Arena {
	termbox.Init()
	x, y := termbox.Size()
	x, y = x / 2, y /2
	a := &Arena {
		FromX : x - (x / 2),
		ToX : x + (x / 2),
		FromY : y - (y / 2),
		ToY : y + (y / 2),
	}
	if cfg != nil {
		a.cfg = *cfg
	} else {
		a.cfg = defaultConfig()
	}
	return a
}

func (a *Arena) Draw(snake []Cell) {
	a.drawBackground()
	a.drawSnake(snake)
	termbox.Flush()
}

func (a *Arena) drawBackground() {
	for i := a.FromX; i < a.ToX; i++ {
		for j := a.FromY; j < a.ToY; j++ {
			termbox.SetCell(i, j, ' ', termbox.ColorDefault, a.cfg.bgColor)
		}
	}
}

func (a *Arena) drawSnake(snake []Cell) {
	for _, cell := range snake {
		termbox.SetCell(cell.X, cell.Y, ' ', a.cfg.snakeColor, termbox.ColorDefault)
	}
}

func (a *Arena) Close() {
	termbox.Close()
}
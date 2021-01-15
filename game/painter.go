package game

import (
	"github.com/nsf/termbox-go"
)

type PainterConfig struct {
	Colors struct {
		Snake termbox.Attribute
		Food  termbox.Attribute
		Bg    termbox.Attribute
	}

	Symbols struct {
		Snake rune
		Food  rune
	}
}

func PainterCfg() *PainterConfig {
	var cfg PainterConfig
	cfg.Colors.Snake = termbox.ColorBlack
	cfg.Colors.Food = termbox.ColorRed
	cfg.Colors.Bg = termbox.ColorGreen

	cfg.Symbols.Snake = ' '
	cfg.Symbols.Food = '●'

	return &cfg
}

type Painter struct {
	cfg   PainterConfig
	arena Arena
}

func NewPainter(cfg *PainterConfig, arena Arena) *Painter {
	if cfg == nil {
		cfg = PainterCfg()
	}

	return &Painter{
		cfg:   *cfg,
		arena: arena,
	}
}

func (p *Painter) Draw(snake *Node, food Cell) {
	for y := p.arena.FromY; y <= p.arena.ToY; y++ {
		for x := p.arena.FromX; x <= p.arena.ToX; x++ {
			termbox.SetBg(x, y, p.cfg.Colors.Bg)
		}
	}

	for snake != nil {
		termbox.SetCell(snake.Coordinate.X, snake.Coordinate.Y, p.cfg.Symbols.Snake, p.cfg.Colors.Snake, p.cfg.Colors.Snake)
		snake = snake.Tail
	}

	termbox.SetCell(food.X, food.Y, p.cfg.Symbols.Food, p.cfg.Colors.Food, p.cfg.Colors.Bg)

	termbox.Flush()
}

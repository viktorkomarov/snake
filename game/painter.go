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

func (p *Painter) Draw(snake *Snake, food Food) {
	for y := p.arena.FromY; y <= p.arena.ToY; y++ {
		for x := p.arena.FromX; x <= p.arena.ToX; x++ {
			termbox.SetBg(x, y, p.cfg.Colors.Bg)
		}
	}

	termbox.Flush()
}

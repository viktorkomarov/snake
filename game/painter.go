package game

import (
	"fmt"
	"time"

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

func NewPainter(cfg *PainterConfig) (*Painter, error) {
	if cfg == nil {
		cfg = PainterCfg()
	}

	if err := termbox.Init(); err != nil {
		return nil, fmt.Errorf("%w: can't init termobox", err)
	}
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	width, height := termbox.Size()
	arena := NewArena(width, height)

	return &Painter{
		cfg:   *cfg,
		arena: arena,
	}, nil
}

func (p *Painter) Close() {
	termbox.Close()
}

func (p *Painter) Draw() {
	for y := p.arena.FromY; y <= p.arena.ToY; y++ {
		for x := p.arena.FromX; x <= p.arena.ToX; x++ {
			termbox.SetBg(x, y, p.cfg.Colors.Bg)
		}
	}
	termbox.Flush()

	time.Sleep(time.Second * 5)
}

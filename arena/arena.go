package arena

import (
	"github.com/nsf/termbox-go"
)

type Config struct {
	snakeColor termbox.Attribute
	foodColor termbox.Attribute
	bgColor termbox.Attribute
}

func defaultConfig() Config {
	return Config {
		snakeColor  : termbox.ColorGreen,
		foodColor   : termbox.ColorRed,
		bgColor     : termbox.ColorBlue,
	}
}

type Arena struct {
	xSize int
	ySize int
	cfg Config
}

func New(cfg *Config) *Arena {
	termbox.Init()
	x, y := termbox.Size()
	x, y = x / 2, y /2
	a := &Arena {
		xSize : x,
		ySize : y,
	}
	if cfg != nil {
		a.cfg = *cfg
	} else {
		a.cfg = defaultConfig()
	}
	return a
}

func (a *Arena) Draw() {
	a.drawBackground()
	termbox.Flush()
}

func (a *Arena) drawBackground() {
	hX, hY := a.xSize / 2, a.ySize / 2
	for i := a.xSize - hX; i < a.xSize + hX; i++ {
		for j := a.ySize - hY; j < a.ySize + hY; j++ {
			termbox.SetCell(i, j, ' ', termbox.ColorDefault, a.cfg.bgColor)
		}
	}
}

func (a *Arena) Close() {
	termbox.Close()
}
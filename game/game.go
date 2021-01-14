package game

import (
	"context"
	"time"
)

type Game struct {
	painter *Painter
	events  <-chan Event
	ticker  *time.Ticker
}

func NewGame(ctx context.Context, painter *Painter, events <-chan Event) *Game {
	return &Game{
		painter: painter,
		ticker:  time.NewTicker(time.Second * 3),
		events:  events,
	}
}

func (g *Game) updateTicker(d time.Duration) {}

func (g *Game) move(lastStep Event) Event {
	select {
	case e := <-g.events:
		return e
	default:
		return lastStep
	}
}

func (g *Game) Start() {
	lastStep := Left
	for _ = range g.ticker.C {
		lastStep = g.move(lastStep)
	}
}

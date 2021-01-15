package game

import (
	"errors"
	"time"
)

var ErrGameInterrupt = errors.New("game was finished by user")

type Game struct {
	painter *Painter
	events  <-chan Event
	ticker  *time.Ticker
	arena   Arena
}

func NewGame(painter *Painter, events <-chan Event, arena Arena) *Game {
	return &Game{
		painter: painter,
		events:  events,
		arena:   arena,
		ticker:  time.NewTicker(time.Millisecond * 50),
	}
}

func (g *Game) updateTicker() {
}

func (g *Game) move(lastStep Event) Event {
	select {
	case e := <-g.events:
		return e
	default:
		return lastStep
	}
}

func (g *Game) Start() (int, error) {
	lastStep := Left

	snake := NewSnake(g.arena, lastStep)
	food := g.arena.RandomCell()

	for _ = range g.ticker.C {
		lastStep = g.move(lastStep)
		if lastStep == Kill {
			return snake.Size(), ErrGameInterrupt
		}

		eat, err := snake.move(lastStep, food)
		if err != nil {
			return snake.Size(), err
		}

		if eat {
			food = g.arena.RandomCell()
		}

		g.painter.Draw(snake.Head, food)
	}

	return 0, nil
}

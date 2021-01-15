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
}

func NewGame(painter *Painter, events <-chan Event) *Game {
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

func (g *Game) Start() (int, error) {
	lastStep := Left

	snake := NewSnake()
	food := NewFood(0, 0, 0, 0)
	for _ = range g.ticker.C {
		event := g.move(lastStep)
		if event == Kill {
			return snake.Score(), ErrGameInterrupt
		}

		err := snake.move(event, food.Coordinate())
		if err != nil {
			return snake.Score(), err
		}

		g.painter.Draw(snake, food)
	}

	return 0, nil
}

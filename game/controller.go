package game

import (
	"context"

	"github.com/nsf/termbox-go"
)

type Event int

const (
	Unknown Event = iota
	Up
	Down
	Left
	Right
	Kill
)

var runeToEvent = map[rune]Event{
	'w': Up, 'W': Up,
	's': Down, 'S': Down,
	'a': Left, 'A': Left,
	'd': Right, 'D': Right,
}

var keyToEvent = map[termbox.Key]Event{
	termbox.KeyArrowDown:  Down,
	termbox.KeyArrowUp:    Up,
	termbox.KeyArrowLeft:  Left,
	termbox.KeyArrowRight: Right,
	termbox.KeyEsc:        Kill, termbox.KeyCtrlC: Kill,
}

type Controller struct {
	events chan Event
}

func NewController(ctx context.Context) *Controller {
	c := &Controller{
		events: make(chan Event),
	}
	go c.run(ctx)

	return c
}

func handleEvent(key termbox.Key, ch rune) Event {
	if event := runeToEvent[ch]; event != Unknown {
		return event
	}

	return keyToEvent[key]
}

func (c *Controller) run(ctx context.Context) {
	output := make(chan Event)
	go func() {
		defer close(output)

		for {
			event := termbox.PollEvent()
			if event.Type == termbox.EventKey { // support only keyboard
				output <- handleEvent(event.Key, event.Ch)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case e := <-output:
			if e != Unknown {
				c.events <- e
			}
		}
	}
}

func (c *Controller) Events() <-chan Event {
	return c.events
}

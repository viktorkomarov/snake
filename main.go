package main

import (
	"time"

	"github.com/nsf/termbox-go"
	"github.com/viktorkomarov/snake/game"
)

func main() {
	arena := game.NewArena(nil)
	snake := game.NewSnake(arena.FromX, arena.ToX, arena.FromY, arena.ToY)
	food := game.NewFood(arena.FromX, arena.ToX, arena.FromY, arena.ToY)

	request := make(chan struct{})
	active := game.UserAction()
	actionByRequest := game.Actions(request, active, snake.NextStep)
	if termbox.IsInit {
		game.Run(time.Millisecond*50, actionByRequest, snake, arena, request, food)
	}
}

package main

import(
	"time"

	"gitlab.com/VictorKomarov/snake/game"
	"github.com/nsf/termbox-go"
)

func main(){
	arena := game.NewArena(nil)
	snake := game.NewSnake(arena.FromX, arena.ToX, arena.FromY, arena.ToY)

	request := make(chan struct{})
	active := game.UserAction()
	actionByRequest := game.Actions(request, active)
	if termbox.IsInit {
		game.Run(time.Second * 1, actionByRequest, snake, arena, request)
	}
}
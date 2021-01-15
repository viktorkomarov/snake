package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nsf/termbox-go"
	"github.com/viktorkomarov/snake/game"
)

func main() {
	err := termboxInit()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	width, height := termbox.Size()
	arena := game.NewArena(width, height)
	painter := game.NewPainter(game.PainterCfg(), arena)
	collector := game.NewController(ctx)
	game := game.NewGame(painter, collector.Events())

	score, err := game.Start()
	fmt.Printf("%d %v\n", score, err)
}

func termboxInit() error {
	err := termbox.Init()
	if err != nil {
		return err
	}

	termbox.SetInputMode(termbox.InputEsc)
	return nil
}

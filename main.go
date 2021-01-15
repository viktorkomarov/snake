package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/viktorkomarov/snake/game"
)

func main() {
	err := initGame()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	width, height := termbox.Size()
	arena := game.NewArena(width, height)
	painter := game.NewPainter(game.PainterCfg(), arena)
	collector := game.NewController(ctx)
	game := game.NewGame(painter, collector.Events(), arena)

	score, err := game.Start()
	termbox.Close()
	fmt.Printf("%d %v\n", score, err)
}

func initGame() error {
	rand.Seed(time.Now().Unix())
	err := termbox.Init()
	if err != nil {
		return err
	}

	termbox.SetInputMode(termbox.InputEsc)
	return nil
}

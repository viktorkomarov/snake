package main

import (
	"context"
	"log"

	"github.com/viktorkomarov/snake/game"
)

func main() {
	painter, err := game.NewPainter(game.PainterCfg())
	if err != nil {
		log.Fatal(err)
	}

	events := game.NewController(context.Background())
	ch := events.Events()
	e := <-ch
	painter.Close()

	log.Printf("%v", e)
}

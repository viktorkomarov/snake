package game

import (
	"time"
	"errors"

	"github.com/nsf/termbox-go"
)

func defineSnakeRoad(key termbox.Key) (NextStep, error) {
	switch key {
	case termbox.KeyArrowUp:
		return Left, nil
	case termbox.KeyArrowDown:
		return Right, nil
	case termbox.KeyArrowLeft:
		return Down, nil
	case termbox.KeyArrowRight:
		return Up, nil
	default:
		return Up, errors.New("unsupported event key")				
	}
}

func UserAction() chan NextStep{
	action := make(chan NextStep)
	go func(){
		for{
			event := termbox.PollEvent()
			if event.Err != nil {
				// impement
			}

			road, err := defineSnakeRoad(event.Key)
			if err != nil {
				// implement
			}
			action <- road
		}
	}()

	return action
}

func Actions(request chan struct{}, changeLastAction chan NextStep, last NextStep) chan NextStep {
	output := make(chan NextStep)

	go func(){
		for{
			select{
			case <-request:
				output <- last
			case update := <- changeLastAction:
				last = update
			}
		}
	}()

	return output
}

func Run(speed time.Duration, action chan NextStep, snake *Snake, arena *Arena, request chan struct{},food Food) {
	arena.Draw(snake.Snapshot(), food.Coordinate())

	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()
	nextTick := time.Now().Add(speed)

	for range ticker.C {
		if time.Now().Before(nextTick) {
			continue
		}
		
		request <- struct{}{}
		step := <- action
		
		snake.MoveByUser(step)

		if eaten := snake.Eat(food.Coordinate()); eaten {
			food.regenerate()
		}

		arena.Draw(snake.Snapshot(), food.Coordinate())
		nextTick = nextTick.Add(speed)
	}
}
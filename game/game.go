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

func Actions(request chan struct{}, changeLastAction chan NextStep) chan NextStep {
	output := make(chan NextStep)

	go func(){
		var last NextStep
		updated := false

		for{
			select{
			case <-request:
				if updated {
					output <- last
				}
			case update := <- changeLastAction:
				last = update
				updated = true	
			}
		}
	}()

	return output
}

func Run(speed time.Duration, action chan NextStep, snake *Snake, arena *Arena, request chan struct{}) {
	
	for{
		request <- struct{}{}
		timer := time.NewTimer(speed)
		select{
		case <- timer.C: // correct ?
			snake.Move()
		case road :=<- action:
			snake.MoveByUser(road)
			<- timer.C
		}

		arena.Draw(snake.Snapshot())
	}
}
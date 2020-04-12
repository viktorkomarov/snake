package main

import (
	"gitlab.com/VictorKomarov/snake/arena"
)

func main(){
	arena := arena.New(nil)
	defer arena.Close()
	arena.Draw()
	for {}
}
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/tictactoe-game-cli/core"
	"github.com/omar0ali/tictactoe-game-cli/game"
)

func main() {
	window := core.CreateWindow("TicTacToe")
	// objects

	// exit channel waiting to get a an exit signal == 0 from either Events() or Update()
	exit := make(chan int)

	// testing a single box
	width, height := window.Screen.Size()
	box := game.CreateBoxHolder(game.Point{X: width / 2, Y: height / 2}, 4)
	// for large box, increament by even numbers: min set to 4

	window.Events(exit,
		func(event tcell.Event) {
			box.InputEvents(event, &window)
		},
	)

	window.Update(exit,
		func(delta float64) {
			// animation goes here
			box.Update(&window)
			box.Draw(&window)
		},
	)

	// exit
	if val := <-exit; val == 0 {
		return
	}
}

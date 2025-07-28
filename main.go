package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/tictactoe-game-cli/core"
	"github.com/omar0ali/tictactoe-game-cli/entities"
	"github.com/omar0ali/tictactoe-game-cli/game"
)

func main() {
	window := core.CreateWindow("TicTacToe")

	// exit channel waiting to get a an exit signal == 0 from either Events() or Update()
	exit := make(chan int)

	// test
	width, height := window.Screen.Size()
	box := entities.CreateBoxHolder(entities.Point{X: width / 2, Y: height / 2}, 4)
	box.SetContent('X')
	box1 := entities.CreateBoxHolder(entities.Point{X: (width / 2) + 5, Y: height / 2}, 4)
	box1.SetContent('O')
	// for large box, increament by even numbers: min set to 4

	// add objects into the game
	gameState := game.GameContext{Window: &window}
	gameState.AddEntity(&box)
	gameState.AddEntity(&box1)

	window.Events(exit,
		func(event tcell.Event) {
			for _, entity := range gameState.GetEntities() {
				entity.InputEvents(event)
			}
		},
	)

	window.Update(exit,
		func(delta float64) {
			// animation goes here
			for _, entity := range gameState.GetEntities() {
				entity.Update(&gameState)
				entity.Draw(&gameState)
			}
		},
	)

	// exit
	if val := <-exit; val == 0 {
		return
	}
}

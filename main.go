package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/tictactoe-game-cli/core"
	"github.com/omar0ali/tictactoe-game-cli/entities"
	"github.com/omar0ali/tictactoe-game-cli/game"
	"github.com/omar0ali/tictactoe-game-cli/views"
)

func main() {
	window := core.CreateWindow("TicTacToe")

	// exit channel waiting to get a an exit signal == 0 from either Events() or Update()
	exit := make(chan int)

	gridView := views.InitGridView(9, 1, 4, 3, &window)
	gameState := game.GameContext{Window: &window, PlayerTurn: game.P1}

	boxes := []*entities.BoxHolder{}
	for i := range gridView.GetItems() {
		box := gridView.GetItems()[i]              // get each box
		boxHolder, ok := box.(*entities.BoxHolder) // try to cast (assertion)
		if !ok {
			panic("Type assertion failed: not a *BoxHolder")
		}
		gameState.AddEntity(boxHolder)   // add it as an entity
		boxes = append(boxes, boxHolder) // this boxes used later | To check for winers
		boxHolder.SetBoxes(&boxes)       // each box should have a ref of all the boxes. Helps to check
		// who wins after each turn
	}

	window.Events(exit,
		func(event tcell.Event) {
			for _, entity := range gameState.GetEntities() {
				entity.InputEvents(event, &gameState)
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

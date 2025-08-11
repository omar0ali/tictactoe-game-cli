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

	gridView := views.InitGridView(9, 1, 4, 3, &window)                         // grid system
	dialog := game.InitDialog(70, game.BottomCenter, window.Screen, "[Window]") // dialog enabled

	logs := game.InitDialog(10, game.TopRight, window.Screen, "[Logs]")
	logs.Log = true
	logs.AddLine("Start Game")

	dialog.AddLine("[TicTacToe Game]")
	dialog.AddLine("* Press 'c' key to close any dialog window. to quit the game 'q' or 'ESC'.")
	dialog.AddLine("* Press 'h' key to to disable dialog dialog window.")
	gameState := game.GameContext{Window: &window, PlayerTurn: game.P1, Dialog: &dialog, Logs: &logs}

	// Add boxes on screen
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

	// Add Dialog on screen.
	gameState.AddEntity(&dialog)
	gameState.AddEntity(&logs)

	window.Events(exit,
		func(event tcell.Event) {
			switch ev := event.(type) {
			case *tcell.EventKey:
				if ev.Rune() == 'r' {
					// restarting the game
					entities.RestartGame(&gameState, &boxes, 0)
					entities.EnableBoxes(&boxes)
				}
			}
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

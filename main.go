package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/tictactoe-game-cli/core"
)

func main() {
	window := core.CreateWindow("TicTacToe")

	// exit channel waiting to get a an exit signal == 0 from either Events() or Update()
	exit := make(chan int)

	window.Events(exit,
		func(event tcell.Event) {
			switch ev := event.(type) {
			case *tcell.EventMouse:
				if ev.Buttons() == tcell.Button1 {
					exit <- 0 // temp: will exit for now
					return
				}
			}
		},
	)

	window.Update(exit,
		func(delta float64) {
			// animation goes here
		},
	)

	// exit
	if val := <-exit; val == 0 {
		return
	}
}

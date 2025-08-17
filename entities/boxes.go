package entities

import (
	"github.com/omar0ali/tictactoe-game-cli/game"
)

var boxes []*BoxHolder

func SetBoxes(listOfBoxes []*BoxHolder) {
	InitBoxes(listOfBoxes)
	boxes = listOfBoxes
}

func InitBoxes(listOfBoxes []*BoxHolder) {
	for _, v := range listOfBoxes {
		v.visible = false
		v.content = ' '
	}
}

func RestartGame(gc *game.GameContext) {
	// rest the game
	gc.PlayerTurn = game.P1
	InitBoxes(boxes)
	if gc.Logs.Log {
		gc.Logs.ClearLines()
		gc.Logs.AddLine("Start New Game")
	}
	EnableBoxes()
}

func GetBoxes() []*BoxHolder {
	if boxes != nil {
		return boxes
	}
	return nil
}

func DisabledBoxes() bool {
	if boxes == nil {
		return false
	}
	for _, box := range boxes {
		box.disable = true
	}
	return true
}

func EnableBoxes() bool {
	if boxes == nil {
		return false
	}
	for _, box := range boxes {
		box.disable = false
	}
	return true
}

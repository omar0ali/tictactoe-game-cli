package entities

import (
	"github.com/omar0ali/tictactoe-game-cli/game"
)

var boxes []*BoxHolder

const (
	MinScore = -2
	MaxScore = 2
)

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

func ColorBoxes(boxes ...*BoxHolder) bool {
	if boxes == nil {
		return false
	}
	for _, box := range boxes {
		box.ChangeStyle()
	}
	return true
}

// Each box should have a ref of all boxes, that is easier to check on very turn, since
// each box has its own inputevents too.
// The function returns if the end of the game reached and the winner.

func IsTerminal() (bool, int) {
	winPatterns := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, pattern := range winPatterns {
		if boxes[pattern[0]].content != ' ' &&
			boxes[pattern[0]].content == boxes[pattern[1]].content &&
			boxes[pattern[1]].content == boxes[pattern[2]].content {
			ColorBoxes(boxes[pattern[0]], boxes[pattern[1]], boxes[pattern[2]])
			// check if O
			if boxes[pattern[0]].content == 'O' {
				return true, 1
			}
			// check if X
			return true, -1
		}
	}

	// check if there are empty boxes
	for i := range boxes {
		if boxes[i].content == ' ' {
			return false, 0
		}
	}
	// no winners - End of the game
	return true, 0
}

func Minimax(isMaximizing bool) int {
	terminal, score := IsTerminal()
	if terminal {
		return score
	}

	if isMaximizing { // O
		best := MinScore
		for i := range boxes {
			if boxes[i].content == ' ' {
				boxes[i].content = 'O'
				val := Minimax(false)
				boxes[i].content = ' ' // undo move
				best = max(val, best)
			}
		}
		return best
	} else { // X
		best := MaxScore
		for i := range boxes {
			if boxes[i].content == ' ' {
				boxes[i].content = 'X'
				val := Minimax(true)
				boxes[i].content = ' ' // undo
				best = min(val, best)
			}
		}
		return best
	}
}

// getAIMove returns the index of the best move for the AI

func GetAIMove() int {
	bestScore := MinScore // min score is 'O' Player 2
	move := -1

	for i := range boxes {
		if boxes[i].content == ' ' {
			boxes[i].content = 'O'
			score := Minimax(false)
			boxes[i].content = ' '
			if score > bestScore {
				bestScore = score
				move = i
			}
		}
	}

	return move
}

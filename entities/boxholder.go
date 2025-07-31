// Package entities
package entities

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/tictactoe-game-cli/game"
	"github.com/omar0ali/tictactoe-game-cli/utils"
)

type BoxHolder struct {
	sPoints utils.Point
	ePoints utils.Point
	content rune
	visible bool
	Boxes   *[]*BoxHolder // each box will have a reference of the board
}

func CreateBoxHolder(startPoint utils.Point, scale int) *BoxHolder {
	minimum := max(scale, 4)
	startX := startPoint.X
	startY := startPoint.Y
	endX := startPoint.X + minimum
	endY := startPoint.Y + minimum - 2
	boxHolder := BoxHolder{
		sPoints: utils.Point{X: startX, Y: startY},
		ePoints: utils.Point{X: endX, Y: endY},
		visible: false,
	}
	return &boxHolder
}

func (b *BoxHolder) GetContent() rune {
	return b.content
}

func (b *BoxHolder) SetBoxes(boxes *[]*BoxHolder) {
	b.Boxes = boxes
}

func (b *BoxHolder) InputEvents(event tcell.Event, gc *game.GameContext) {
	switch ev := event.(type) {
	case *tcell.EventMouse:
		if ev.Buttons() == tcell.Button1 {
			if b.visible {
				return
			}
			mouseX, mouseY := ev.Position()
			if mouseX >= b.sPoints.X && mouseY >= b.sPoints.Y {
				if mouseX <= b.ePoints.X && mouseY <= b.ePoints.Y {
					switch gc.PlayerTurn {
					case game.P1:
						b.content = 'X'
						gc.PlayerTurn = game.P2
					case game.P2:
						b.content = 'O'
						gc.PlayerTurn = game.P1
					}
					// and
					b.visible = !b.visible

					// needs refactor
					if b.TicTacToePatternsNew(b.content) { // if win reset the game.
						fmt.Printf("Win: %c", b.content)
					}
				}
			}
		}
	}
}

func (b *BoxHolder) GetTopLeftCoords() utils.Point {
	return b.sPoints
}

func (b *BoxHolder) GetBottomRightCoords() utils.Point {
	return b.ePoints
}

func (b *BoxHolder) GetTopRightCoords() utils.Point {
	return utils.Point{
		X: b.ePoints.X,
		Y: b.sPoints.Y,
	}
}

func (b *BoxHolder) GetBottomLeftCoords() utils.Point {
	return utils.Point{
		X: b.sPoints.X,
		Y: b.ePoints.Y,
	}
}

func (b *BoxHolder) GetBoxHeight() int {
	return b.ePoints.Y - b.sPoints.Y
}

func (b *BoxHolder) GetBoxWidth() int {
	return b.ePoints.X - b.sPoints.X
}

func (b *BoxHolder) Update(gs *game.GameContext) {
	if b.visible {
		middleX := (b.GetBoxWidth() / 2) + b.sPoints.X
		middleY := (b.GetBoxHeight() / 2) + b.sPoints.Y
		gs.Window.SetContent(middleX, middleY, b.content)
	}
}

func (b *BoxHolder) Draw(gs *game.GameContext) {
	// draw corners
	gs.Window.SetContent(b.GetTopLeftCoords().X, b.GetTopLeftCoords().Y, tcell.RuneULCorner)
	gs.Window.SetContent(b.GetTopRightCoords().X, b.GetTopRightCoords().Y, tcell.RuneURCorner)
	gs.Window.SetContent(b.GetBottomLeftCoords().X, b.GetBottomLeftCoords().Y, tcell.RuneLLCorner)
	gs.Window.SetContent(b.GetBottomRightCoords().X, b.GetBottomRightCoords().Y, tcell.RuneLRCorner)
	// draw lines
	for i := 1; i < b.GetBoxHeight(); i++ {
		gs.Window.SetContent(b.GetTopLeftCoords().X, b.GetTopLeftCoords().Y+i, tcell.RuneVLine)
	}
	for i := 1; i < b.GetBoxHeight(); i++ {
		gs.Window.SetContent(b.GetTopRightCoords().X, b.GetTopRightCoords().Y+i, tcell.RuneVLine)
	}
	for i := 1; i < b.GetBoxWidth(); i++ {
		gs.Window.SetContent(b.GetTopLeftCoords().X+i, b.GetTopLeftCoords().Y, tcell.RuneHLine)
	}
	for i := 1; i < b.GetBoxWidth(); i++ {
		gs.Window.SetContent(b.GetBottomLeftCoords().X+i, b.GetBottomLeftCoords().Y, tcell.RuneHLine)
	}
}

// new version
//
// Each box should have a ref of all boxes, that is easier to check on very turn, since
// each box has its own inputevents too.

func (b *BoxHolder) TicTacToePatternsNew(content rune) bool {
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

	board := *b.Boxes
	for _, pattern := range winPatterns {
		if board[pattern[0]].content == content &&
			board[pattern[1]].content == content &&
			board[pattern[2]].content == content {
			return true
		}
	}

	return false
}

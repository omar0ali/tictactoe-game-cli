// Package entities
package entities

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/tictactoe-game-cli/game"
	"github.com/omar0ali/tictactoe-game-cli/utils"
)

var id int = 0

type BoxHolder struct {
	ID      int
	sPoints utils.Point
	ePoints utils.Point
	content rune
	visible bool
	disable bool
}

func CreateBoxHolder(startPoint utils.Point, scale int) *BoxHolder {
	minimum := max(scale, 4)
	startX := startPoint.X
	startY := startPoint.Y
	endX := startPoint.X + minimum
	endY := startPoint.Y + minimum - 2
	boxHolder := BoxHolder{
		ID:      id,
		sPoints: utils.Point{X: startX, Y: startY},
		ePoints: utils.Point{X: endX, Y: endY},
		visible: false,
		disable: false,
	}
	id++
	return &boxHolder
}

func (b *BoxHolder) GetContent() rune {
	return b.content
}

func CheckGame(gc *game.GameContext, b *BoxHolder) bool {
	status, win := IsTerminal()
	gc.Logs.AddLine(fmt.Sprintf("Place: %c | Box: %d", b.content, b.ID))
	if status {
		if win != 0 {
			gc.Logs.AddLine(fmt.Sprintf("* Winner: %c", b.content))
			DisabledBoxes()
		} else {
			gc.Logs.AddLine("* Draw")
		}
		gc.Dialog.ClearLines()
		gc.Dialog.AddLine("* You can press 'r' key to restart the game at any time.")
		gc.Dialog.AddLine("* Press 'q' to quit.")
		gc.Dialog.SetVisible(true)
		return status
	}
	gc.Logs.AddLine("--------------")
	if gc.PlayerTurn == game.P1 {
		gc.Logs.AddLine("Turn: Player 1")
	} else {
		gc.Logs.AddLine("Turn: Player 2 or AI")
	}
	return status
}

func (b *BoxHolder) SwitchTurn(gc *game.GameContext) {
	if gc.PlayerTurn == game.P1 {
		b.SetContent('X')
		gc.PlayerTurn = game.P2
	} else {
		b.SetContent('O')
		gc.PlayerTurn = game.P1
	}

	if CheckGame(gc, b) {
		return
	}
}

func (b *BoxHolder) InputEvents(event tcell.Event, gc *game.GameContext) {
	switch ev := event.(type) {
	case *tcell.EventMouse:
		if ev.Buttons() == tcell.Button1 {
			if b.disable {
				return
			}
			if b.visible {
				return
			}
			mouseX, mouseY := ev.Position()
			if mouseX >= b.sPoints.X && mouseY >= b.sPoints.Y {
				if mouseX <= b.ePoints.X && mouseY <= b.ePoints.Y {
					b.SwitchTurn(gc)
				}
			}
		}
	case *tcell.EventKey:
		if ev.Rune() == 'a' {
			if gc.PlayerTurn != game.P2 {
				return
			}
			getAiMove := GetAIMove()
			boxes[getAiMove].SetContent('O')
			gc.PlayerTurn = game.P1
			if CheckGame(gc, boxes[getAiMove]) {
				return
			}
		}
	}
}

func (b *BoxHolder) GetWinningPlayer(content rune) int {
	if content == 'X' {
		return int(game.P1)
	}
	return int(game.P2)
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

func (b *BoxHolder) SetContent(content rune) {
	b.content = content
	b.visible = true
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

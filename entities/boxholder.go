// Package entities
package entities

import (
	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/tictactoe-game-cli/game"
	"github.com/omar0ali/tictactoe-game-cli/utils"
)

type BoxHolder struct {
	sPoints utils.Point
	ePoints utils.Point
	content rune
	visible bool
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

func (b *BoxHolder) InputEvents(event tcell.Event) {
	switch ev := event.(type) {
	case *tcell.EventMouse:
		if ev.Buttons() == tcell.Button1 {
			mouseX, mouseY := ev.Position()
			if mouseX >= b.sPoints.X && mouseY >= b.sPoints.Y {
				if mouseX <= b.ePoints.X && mouseY <= b.ePoints.Y {
					b.visible = !b.visible
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

func (b *BoxHolder) SetContent(rune rune) {
	b.content = rune
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

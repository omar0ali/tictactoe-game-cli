// Package entities
package entities

import (
	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/tictactoe-game-cli/game"
)

type Point struct {
	X, Y int
}
type boxHolder struct {
	sPoints Point
	ePoints Point
	content rune
	Visible bool
}

func CreateBoxHolder(startPoint Point, scale int) boxHolder {
	minimum := max(scale, 4)
	startX := startPoint.X
	startY := startPoint.Y
	endX := startPoint.X + minimum
	endY := startPoint.Y + minimum - 2
	boxHolder := boxHolder{
		sPoints: Point{startX, startY},
		ePoints: Point{endX, endY},
		Visible: false,
	}
	return boxHolder
}

func (b *boxHolder) InputEvents(event tcell.Event) {
	switch ev := event.(type) {
	case *tcell.EventMouse:
		if ev.Buttons() == tcell.Button1 {
			mouseX, mouseY := ev.Position()
			if mouseX >= b.sPoints.X && mouseY >= b.sPoints.Y {
				if mouseX <= b.ePoints.X && mouseY <= b.ePoints.Y {
					b.Visible = !b.Visible
				}
			}
		}
	}
}

func (b *boxHolder) GetTopLeftCoords() Point {
	return b.sPoints
}

func (b *boxHolder) GetBottomRightCoords() Point {
	return b.ePoints
}

func (b *boxHolder) GetTopRightCoords() Point {
	return Point{
		X: b.ePoints.X,
		Y: b.sPoints.Y,
	}
}

func (b *boxHolder) GetBottomLeftCoords() Point {
	return Point{
		X: b.sPoints.X,
		Y: b.ePoints.Y,
	}
}

func (b *boxHolder) GetBoxHeight() int {
	return b.ePoints.Y - b.sPoints.Y
}

func (b *boxHolder) GetBoxWidth() int {
	return b.ePoints.X - b.sPoints.X
}

func (b *boxHolder) SetContent(rune rune) {
	b.content = rune
}

func (b *boxHolder) Update(gs *game.GameContext) {
	if b.Visible {
		middleX := (b.GetBoxWidth() / 2) + b.sPoints.X
		middleY := (b.GetBoxHeight() / 2) + b.sPoints.Y
		gs.Window.SetContent(middleX, middleY, b.content)
	}
}

func (b *boxHolder) Draw(gs *game.GameContext) {
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

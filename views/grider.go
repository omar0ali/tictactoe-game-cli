// Package views
package views

import (
	"github.com/omar0ali/tictactoe-game-cli/core"
	"github.com/omar0ali/tictactoe-game-cli/entities"
	"github.com/omar0ali/tictactoe-game-cli/utils"
)

type Box interface {
	GetTopLeftCoords() utils.Point
	GetBottomRightCoords() utils.Point
	GetTopRightCoords() utils.Point
	GetBottomLeftCoords() utils.Point
	GetBoxHeight() int
	GetBoxWidth() int
}

type GridView struct {
	items []Box
}

func InitGridView(no, gap, scale, max int, window *core.Window) GridView {
	width, height := window.Screen.Size()

	// calculate the total width to use
	totalBoxesWidth := max * scale
	totalGapsWidth := (max - 1) * gap
	totalRowWidth := totalBoxesWidth + totalGapsWidth

	startX := (width - totalRowWidth) / 2 // this will give us the correct coords center of screen

	boxes := []Box{}
	currentXPos := 1
	currentYPos := (height / 2) - ((no / max) + (scale - gap))
	tempStartX := startX
	for range no {
		if currentXPos > max {
			currentYPos += scale
			currentXPos = 1
			tempStartX = startX
		}
		box := entities.CreateBoxHolder(utils.Point{X: tempStartX, Y: currentYPos}, scale)
		box.SetContent('X') // TODO: Update as needed
		boxes = append(boxes, box)
		tempStartX = tempStartX + scale + gap
		currentXPos += 1

	}

	return GridView{items: boxes}
}

func (g *GridView) GetItems() []Box {
	return g.items
}

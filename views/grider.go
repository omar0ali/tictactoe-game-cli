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

func InitGridView(no, gap, scale int, window *core.Window) GridView {
	// we know that each has the same size width and height
	width, height := window.Screen.Size()

	// calculate the total width to use
	totalBoxesWidth := no * scale
	totalGapsWidth := (no - 1) * gap
	totalRowWidth := totalBoxesWidth + totalGapsWidth

	startX := (width - totalRowWidth) / 2 // this will give us the correct coords center of screen

	boxes := []Box{}
	for range no {
		box := entities.CreateBoxHolder(utils.Point{X: startX, Y: height / 2}, scale)
		box.SetContent('X') // TODO: Update as needed
		boxes = append(boxes, box)
		startX = startX + scale + gap
	}

	return GridView{items: boxes}
}

func (g *GridView) GetItems() []Box {
	return g.items
}

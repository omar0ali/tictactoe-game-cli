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
	boxes := []Box{}

	// total width
	columns := max * scale
	gaps := (max - 1) * gap
	totalRowWidth := columns + gaps
	startX := (width - totalRowWidth) / 2

	// total height

	// this one made sense lol | we always have one row missing when there is a fraction in result
	rows := no / max
	if no%max != 0 {
		rows++
	}
	totalHeight := rows*scale + (rows-1)*gap // we always have 1 gap less of the total rows
	startY := (height - totalHeight) / 2

	currentCol := startX
	currentRow := startY

	col := 0
	for range no {
		box := entities.CreateBoxHolder(utils.Point{X: currentCol, Y: currentRow}, scale)
		box.SetContent('X')
		boxes = append(boxes, box)
		col++ // added one
		if col == max {
			// move to next row
			col = 0
			currentCol = startX
			currentRow += scale + gap
		} else {
			// move to next column
			currentCol += scale + gap
		}
	}

	return GridView{items: boxes}
}

func (g *GridView) GetItems() []Box {
	return g.items
}

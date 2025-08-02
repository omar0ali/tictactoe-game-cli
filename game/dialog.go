package game

import (
	"github.com/gdamore/tcell/v2"
)

type Dialog struct {
	startX  int
	endX    int
	lines   []string
	visible bool
}

func InitDialog(maxWidth int, screen tcell.Screen) Dialog {
	minWidth := max(maxWidth, 40)
	width, _ := screen.Size()
	sX := (width / 2) - (minWidth / 2)
	dialog := Dialog{
		startX:  sX,
		endX:    sX + minWidth,
		visible: true,
	}
	return dialog
}

func (d *Dialog) AddLine(text string) {
	maxWidth := d.endX - d.startX - 1
	runes := []rune(text)

	for i := 0; i < len(runes); i += maxWidth {
		end := min(i+maxWidth, len(runes))
		d.lines = append(d.lines, string(runes[i:end]))
	}
}

func (d *Dialog) SetVisible(visible bool) {
	d.visible = visible
}

func (d *Dialog) ClearLines() {
	d.lines = d.lines[:0]
	// clear(d.lines) // problem with this, it just removes cotnent of each element in the slice
	// But keeps Same len/cap, and start the append not starting from 0.
}

func (d *Dialog) Update(gs *GameContext) {}

func (d *Dialog) Draw(gs *GameContext) {
	if !d.visible {
		return
	}

	screen := gs.Window.Screen
	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorWhite)
	_, height := screen.Size()
	height = height / 2
	maxHeight := len(d.lines)

	// corners
	screen.SetContent(d.startX, height/2, tcell.RuneULCorner, nil, style)
	screen.SetContent(d.endX, height/2, tcell.RuneURCorner, nil, style)
	screen.SetContent(d.startX, maxHeight+(height/2)+1, tcell.RuneLLCorner, nil, style)
	screen.SetContent(d.endX, maxHeight+(height/2)+1, tcell.RuneLRCorner, nil, style)

	// top line
	distance := d.endX - d.startX - 1
	for i := range distance {
		screen.SetContent(d.startX+i+1, height/2, tcell.RuneHLine, nil, style)
	}

	// bottom line
	for i := range distance {
		screen.SetContent(d.startX+i+1, maxHeight+(height/2)+1, tcell.RuneHLine, nil, style)
	}

	// left line
	for i := range maxHeight {
		screen.SetContent(d.startX, (height/2)+i+1, tcell.RuneVLine, nil, style)
	}
	// right line
	for i := range maxHeight {
		screen.SetContent(d.endX, (height/2)+i+1, tcell.RuneVLine, nil, style)
	}

	// Draw Text
	for y, line := range d.lines {
		runes := []rune(line)
		x := 0
		for ; x < len(runes); x++ {
			screen.SetContent(d.startX+x+1, (height/2)+y+1, runes[x], nil, style)
		}
		// this will just fill empy spaces on the line to cover or hide whats behind it.
		for ; x < distance; x++ {
			screen.SetContent(d.startX+x+1, (height/2)+y+1, ' ', nil, style)
		}
	}
	for x, v := range []rune("To close window press 'c'.") {
		screen.SetContent(d.startX+x+1, maxHeight+(height/2)+2, v, nil, style)
	}
}

func (d *Dialog) InputEvents(event tcell.Event, gc *GameContext) {
	switch ev := event.(type) {
	case *tcell.EventKey:
		if ev.Rune() == 'c' {
			d.visible = false
			d.ClearLines()
		}
	}
}

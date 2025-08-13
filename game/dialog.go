package game

import (
	"github.com/gdamore/tcell/v2"
)

const MaxLogsHeight int = 10

type Distance struct {
	StartX, EndX int
}

func (d *Distance) GetMaxWidth() int {
	return d.EndX - d.StartX - 1
}

type Position int

const (
	BottomLeft Position = iota
	BottomRight
	BottomCenter
	TopLeft
	TopRight
	TopCenter
)

type Dialog struct {
	Distance Distance
	Position Position
	lines    []string
	visible  bool
	Log      bool
	screen   *tcell.Screen
	title    string
}

func InitDialog(maxWidth int, position Position, screen tcell.Screen, title string) Dialog {
	minWidth := max(maxWidth, 40)
	width, _ := screen.Size()
	var sX int
	switch position {
	case BottomCenter, TopCenter:
		sX = (width / 2) - (minWidth / 2)
	case BottomLeft, TopLeft:
		sX = 0
	case BottomRight, TopRight:
		sX = width - minWidth - 1
	}

	return Dialog{
		Distance: Distance{
			sX, sX + minWidth,
		},
		visible:  true,
		Position: position,
		screen:   &screen,
		title:    title,
	}
}

func (d *Dialog) AddLine(text string) {
	maxWidth := d.Distance.GetMaxWidth()
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

func (d *Dialog) GetScreenHightPosition() (int, int) {
	_, screenHeight := (*d.screen).Size()
	maxHeight := min(len(d.lines), MaxLogsHeight)
	height := 0
	switch d.Position {
	case TopCenter, TopRight, TopLeft:
	case BottomCenter, BottomLeft, BottomRight:
		height = screenHeight - maxHeight - 3
	}
	return height, maxHeight
}

func (d *Dialog) Draw(gs *GameContext) {
	if !d.visible {
		return
	}
	screen := gs.Window.Screen
	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorWhite)
	height, maxHeight := d.GetScreenHightPosition()

	// corners
	screen.SetContent(d.Distance.StartX, height, tcell.RuneULCorner, nil, style)
	screen.SetContent(d.Distance.EndX, height, tcell.RuneURCorner, nil, style)
	screen.SetContent(d.Distance.StartX, maxHeight+height+1, tcell.RuneLLCorner, nil, style)
	screen.SetContent(d.Distance.EndX, maxHeight+(height)+1, tcell.RuneLRCorner, nil, style)

	// top line
	// need to add the title
	distance := d.Distance.GetMaxWidth()
	title := []rune(d.title)
	for i := range distance {
		if i >= 0 && i < len(title)+1 {
			if i == 0 { // opening bracket for the title
				screen.SetContent(d.Distance.StartX+i+1, height, '[', nil, style)
				continue
			}
			screen.SetContent(d.Distance.StartX+i+1, height, title[i-1], nil, style)
			continue
		}
		if i == len(title)+1 { // closing bracket for the title
			screen.SetContent(d.Distance.StartX+i+1, height, ']', nil, style)
			continue
		}
		screen.SetContent(d.Distance.StartX+i+1, height, tcell.RuneHLine, nil, style)
	}

	// bottom line
	for i := range distance {
		screen.SetContent(d.Distance.StartX+i+1, maxHeight+(height)+1, tcell.RuneHLine, nil, style)
	}

	// left line
	for i := range maxHeight {
		screen.SetContent(d.Distance.StartX, (height)+i+1, tcell.RuneVLine, nil, style)
	}
	// right line
	for i := range maxHeight {
		screen.SetContent(d.Distance.EndX, (height)+i+1, tcell.RuneVLine, nil, style)
	}

	// Draw Text
	start := 0
	if d.Log {
		if len(d.lines) > MaxLogsHeight {
			start = len(d.lines) - MaxLogsHeight
		}
	}
	for y, line := range d.lines[start:] {
		runes := []rune(line)
		x := 0
		for ; x < len(runes); x++ {
			screen.SetContent(d.Distance.StartX+x+1, (height)+y+1, runes[x], nil, style)
		}
		// this will just fill empy spaces on the line to cover or hide whats behind it.
		for ; x < distance; x++ {
			screen.SetContent(d.Distance.StartX+x+1, (height)+y+1, ' ', nil, style)
		}
	}
	if d.Log {
		return
	}
	x := 0
	displayText := []rune("[c] Close Window")
	for ; x < len(displayText); x++ {
		screen.SetContent(d.Distance.StartX+x+1, maxHeight+(height)+2, displayText[x], nil, style)
	}
	for ; x < distance; x++ {
		screen.SetContent(d.Distance.StartX+x+1, maxHeight+(height)+2, ' ', nil, style)
	}
}

func (d *Dialog) InputEvents(event tcell.Event, gc *GameContext) {
	switch ev := event.(type) {
	case *tcell.EventKey:
		if d.Log {
			return
		}
		if ev.Rune() == 'c' {
			d.visible = false
			d.ClearLines()
		}
		// remove dialog window from the game
		if ev.Rune() == 'h' {
			gc.RemoveEntity(d)
		}
	}
}

func (d *Dialog) IsVisible() bool {
	return d.visible
}

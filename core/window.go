// Package core
package core

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Screen tcell.Screen
	Style  tcell.Style
}

func CreateWindow(title string) Window {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	err = screen.Init()
	if err != nil {
		panic(err)
	}
	screen.SetTitle(title)
	screen.EnableMouse()
	return Window{
		Screen: screen,
		Style:  tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorGreenYellow),
	}
}

func (s *Window) Events(ListenKeyEvents func(tcell.Event)) {
	for {
		event := s.Screen.PollEvent()
		switch ev := event.(type) {
		case *tcell.EventResize:
			s.Screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyESC || ev.Rune() == 'q' {
				s.Close()
			}
		}
		ListenKeyEvents(event)
	}
}

func (s *Window) Update(
	ListenForUpdates func(delta float64),
) {
	ticker := time.NewTicker(33 * time.Millisecond)
	defer ticker.Stop()

	var delta float64
	last := time.Now()
	for range ticker.C {
		now := time.Now()
		delta = now.Sub(last).Seconds()
		last = now

		s.Screen.Clear()

		lenStr := []rune(fmt.Sprintf("FPS: %.2f", (1 / delta)))
		for i, r := range lenStr {
			s.SetContent(i, 0, r)
		}

		ListenForUpdates(delta)

		s.Screen.Show()
	}
}

func (s *Window) Close() {
	s.Screen.Fini()
	os.Exit(0)
}

func (s *Window) SetContent(x, y int, prune rune) {
	s.Screen.SetContent(x, y, prune, nil, s.Style)
}

func (s *Window) SetContentWithStyle(x, y int, prune rune, style tcell.Style) {
	s.Screen.SetContent(x, y, prune, nil, style)
}

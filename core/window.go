// Package core
package core

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Screen tcell.Screen
	Ticker *time.Ticker
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
		Ticker: time.NewTicker(33 * time.Millisecond),
		Style:  tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorGreenYellow),
	}
}

func (s *Window) Events(
	exit chan int,
	ListenKeyEvents func(tcell.Event),
) {
	go func() {
		for {
			event := s.Screen.PollEvent()
			switch ev := event.(type) {
			case *tcell.EventResize:
				s.Screen.Clear()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyESC || ev.Rune() == 'q' {
					exit <- 0
					s.Close()
					return
				}
			}
			ListenKeyEvents(event)
		}
	}()
}

func (s *Window) Update(
	exit chan int,
	ListenForUpdates func(delta float64),
) {
	var delta float64
	go func() {
		last := time.Now()
		for {
			select {
			case <-s.Ticker.C:
				now := time.Now()
				delta = now.Sub(last).Seconds()
				last = now

				s.Screen.Clear()

				lenStr := []rune(fmt.Sprintf("Delta: %.4f s | FPS: %.2f", delta, (1 / delta)))
				for i, r := range lenStr {
					s.SetContent(i, 0, r)
				}

				ListenForUpdates(delta)

				s.Screen.Show()

			case val := <-exit:
				if val == 0 {
					s.Close()
					return
				}
			}
		}
	}()
}

func (s *Window) Close() {
	s.Ticker.Stop()
	s.Screen.Fini()
}

func (s *Window) SetContent(x, y int, prune rune) {
	s.Screen.SetContent(x, y, prune, nil, s.Style)
}

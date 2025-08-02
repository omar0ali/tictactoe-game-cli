// Package game
package game

import (
	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/tictactoe-game-cli/core"
)

type PlayerTurn int

const (
	P1 PlayerTurn = 1
	P2 PlayerTurn = 2
)

type (
	Entity interface {
		Draw(gs *GameContext)
		Update(gs *GameContext)
		InputEvents(event tcell.Event, gc *GameContext)
	}
	GameContext struct {
		Window     *core.Window
		entities   []Entity
		PlayerTurn PlayerTurn
		Dialog     Dialog
	}
)

func (gs *GameContext) AddEntity(entity ...Entity) {
	gs.entities = append(gs.entities, entity...)
}

func (gs *GameContext) GetEntities() []Entity {
	return gs.entities
}

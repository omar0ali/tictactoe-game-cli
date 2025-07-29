// Package game
package game

import (
	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/tictactoe-game-cli/core"
)

type (
	Entity interface {
		Draw(gs *GameContext)
		Update(gs *GameContext)
		InputEvents(event tcell.Event)
	}
	GameContext struct {
		Window   *core.Window
		entities []Entity
	}
)

func (gs *GameContext) AddEntity(entity ...Entity) {
	gs.entities = append(gs.entities, entity...)
}

func (gs *GameContext) GetEntities() []Entity {
	return gs.entities
}

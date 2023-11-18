package game

import (
	"github.com/Scullder/poker/internal/deck"
	"github.com/Scullder/poker/internal/player"
)

type Game struct {
	Id      int
	Players []player.Player
	Deck    deck.Deck
}

func MakeGame(id int) Game {
	return Game{
		Id:      id,
		Players: make([]player.Player, 0, 10),
		Deck:    deck.MakeDeck(),
	}
}

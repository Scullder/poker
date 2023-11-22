package game

import (
	"errors"

	"github.com/Scullder/poker/internal/deck"
	"github.com/Scullder/poker/internal/player"
)

type GameInterface interface {
	AddPlayer(player player.Player)
	RemovePlayer(id int)
	FindPlayer(id int) player.Player
	SetDeck(deck deck.DeckInterface)
	PreFlop()
	Flop()
	Turn()
	River()
}

type Game struct {
	Id            int
	Players       []player.Player
	ActivePlayers []*player.Player
	Deck          deck.DeckInterface
	Table         []deck.Card
}

func NewGame(id int, deck deck.DeckInterface) *Game {
	return &Game{
		Id:      id,
		Players: make([]player.Player, 0, 10),
		Deck:    deck,
	}
}

func (g *Game) AddPlayer(player player.Player) {
	g.Players = append(g.Players, player)
}

func (g *Game) RemovePlayer(id int) error {
	for i, player := range g.Players {
		if player.Id == id {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
			return nil
		}
	}

	return errors.New("player not found")
}

func (g *Game) SetDeck(deck deck.DeckInterface) {
	g.Deck = deck
}

func (g *Game) PreFlop() {

}

func (g *Game) Flop() {

}

func (g *Game) Turn() {

}

func (g *Game) River() {

}

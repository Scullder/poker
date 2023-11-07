package player

/*
* TODO: Observer for game actions. EvalCombo for every game state(preflop, flop, turn, river)
 */

import (
	"errors"
	"fmt"

	"github.com/Scullder/poker/internal/deck"
)

type PlayerInterface interface {
	SetHand([]deck.Card)
}

type Player struct {
	Name         string
	Hand         []deck.Card
	CurrentCombo []deck.Deck
}

func MakePlayer(name string) Player {
	return Player{
		Name: name,
		Hand: make([]deck.Card, 0, 2),
	}
}

func (p *Player) SetHand(cards []deck.Card) {
	p.Hand = cards[:2]
}

// Hand + Table
func (p Player) EvalCombo(table []deck.Card) (int, error) {
	if len(p.Hand) != 2 {
		return 0, errors.New("Player don't have any cards")
	}

	combo := append(p.Hand[:], table...)
	fmt.Println(combo)

	return 1, nil
}

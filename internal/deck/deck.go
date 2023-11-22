package deck

import (
	"math/rand"
)

type DeckInterface interface {
	Shuffle()
	DrawCards(count int) []Card
	Restore()
}

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	cards := generateCards()
	deck := &Deck{Cards: cards}
	deck.Shuffle()

	return deck
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len((*d).Cards), func(i, j int) {
		(*d).Cards[i], (*d).Cards[j] = (*d).Cards[j], (*d).Cards[i]
	})
}

func (d *Deck) DrawCards(count int) []Card {
	last := len((*d).Cards) - 1
	(*d).Cards = (*d).Cards[last-count:]

	return (*d).Cards
}

func (d *Deck) Restore() {
	d.Cards = generateCards()
	d.Shuffle()
}

func generateCards() []Card {
	var (
		nSuits  = 4
		nValues = 13
		cards   = make([]Card, 52)
	)

	n := 0
	for i := 1; i <= nSuits; i++ {
		for j := 1; j <= nValues; j++ {
			cards[n] = Card{Suit: i, Val: j}
			n++
		}
	}

	return cards
}

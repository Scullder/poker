package deck

import (
	"math/rand"
)

type DeckInterface interface {
	Shuffle()
	Take(count int) []Card
}

type Deck struct {
	Cards []Card
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

func MakeDeck() Deck {
	cards := generateCards()
	deck := Deck{Cards: cards}
	deck.Shuffle()

	return deck
}

// Deck behavior
func (d *Deck) Shuffle() {
	rand.Shuffle(len((*d).Cards), func(i, j int) {
		(*d).Cards[i], (*d).Cards[j] = (*d).Cards[j], (*d).Cards[i]
	})
}

func (d *Deck) Take(count int) []Card {
	last := len((*d).Cards) - 1
	(*d).Cards = (*d).Cards[last-count:]

	return (*d).Cards
}

func StringSuite(num int) string {
	suites := [4]string{
		"hearts",   // чирвы
		"spades",   // пики
		"diamonds", // бубны
		"clubs",    // трефы
	}

	return suites[num]
}

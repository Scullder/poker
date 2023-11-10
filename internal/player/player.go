package player

/*
* TODO: Observer for game actions. EvalCombo for every game state(preflop, flop, turn, river)
* TODO: Ace - 1 and 13 at same time
 */

import (
	"errors"
	"fmt"
	"sort"

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

func (p Player) EvalCombo(table []deck.Card) (int, error) {
	if len(p.Hand) != 2 {
		return 0, errors.New("Player don't have any cards")
	}

	// Hand + Table
	cards := append(p.Hand[:], table...)

	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].Val > cards[j].Val
	})
	fmt.Println(cards)

	// representetion of the deck, Array: 13(value) x 4(suit)
	deckValues := countValues(cards)

	combinations := map[string][]deck.Card{
		"royal flush":    {},
		"straight flush": {},
		"four":           {},
		"full house":     {},
		"flush":          {},
		"straight":       {},
		"three":          {},
		"two pair":       {},
		"pair":           {},
		"high card":      {},
	}

	// counter for straight combinations
	straight := [4]int{}
	// counter and key for "pairs" combinations
	suitIndexes := []int{}
	key := ""

	for i := len(deckValues) - 1; i > 0; i-- {
		for j, suit := range deckValues[i] {
			straight[j] += suit
			if suit != 0 {
				suitIndexes = append(suitIndexes, j)
			}
		}

		if len(suitIndexes) == 2 {
			key = "pair"
		} else if len(suitIndexes) == 3 {
			key = "three"
		} else if len(suitIndexes) == 4 {
			key = "four"
		}

		if key == "pair" && len(combinations["pair"]) != 0 {
			key = "two pair"
		} else if (key == "pair" && len(combinations["three"]) != 0) || (key == "three" && len(combinations["pair"]) != 0) {
			key = "full house"
		}

		if len(combinations[key]) == 0 && key != "" {
			for _, suitIndex := range suitIndexes {
				combinations[key] = append(combinations[key], deck.Card{Val: i, Suit: suitIndex})
			}
		}

		suitIndexes = []int{}
		key = ""
	}

	fmt.Println(combinations)

	return 1, nil
}

func countValues(cards []deck.Card) [14][4]int {
	// just ignore a 0 index(1 is Ace)
	values := [14][4]int{}

	for _, card := range cards {
		values[card.Val][card.Suit] = 1
		//fmt.Println(card)
	}

	return values
}

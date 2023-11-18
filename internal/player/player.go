package player

/*
* TODO: Observer for game actions. EvalCombo for every game state(preflop, flop, turn, river)
* TODO: Ace - 1 and 13 at same time
 */

import (
	"errors"
	"sort"

	"github.com/Scullder/poker/internal/deck"
)

type PlayerInterface interface {
	SetHand([]deck.Card)
}

type Player struct {
	Name         string
	Id           int
	Hand         []deck.Card
	CurrentCombo []deck.Deck
}

func MakePlayer(name string, id int) Player {
	return Player{
		Name: name,
		Hand: make([]deck.Card, 0, 2),
	}
}

func (p *Player) SetHand(cards []deck.Card) {
	p.Hand = cards[:2]
}

func (p Player) EvalCombo(table []deck.Card) (string, []deck.Card, error) {
	if len(p.Hand) != 2 {
		return "", nil, errors.New("Player don't have any cards")
	}

	// Hand + Table
	cards := append(p.Hand[:], table...)

	key, cards := EvalCombo(cards)

	return key, cards, nil
}

func EvalCombo(cards []deck.Card) (string, []deck.Card) {
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].Val > cards[j].Val
	})

	combinations := map[string][]deck.Card{
		"royal flush":    {},
		"straight flush": {},
		"four":           {},
		"full house":     {},
		"flush":          {},
		"straight":       {},
		"three":          {},
		"two pairs":      {},
		"pair":           {},
		"high card":      {},
	}

	combinationsPriority := []string{
		"royal flush",
		"straight flush",
		"four",
		"full house",
		"flush",
		"straight",
		"three",
		"two pairs",
		"pair",
		"high card",
	}

	// representetion of the deck, Array: 14(value) x 4(suit)
	deckValues := [14][4]int{}

	for _, card := range cards {
		deckValues[card.Val][card.Suit] = 1
	}

	// count not null suits for value[i]
	suitIndexes := []int{}
	key := ""

	// straight counter
	straightBuilder := []deck.Card{}

	// straight(+royal) flush counter
	straightFlushBuilder := [4][]deck.Card{}

	for i, suit := range deckValues[1] {
		if suit == 1 {
			straightFlushBuilder[i] = append(straightFlushBuilder[i], deck.Card{Val: i, Suit: i})
		}
	}

	out := false

	// flush
	flushBuilder := [4][]deck.Card{}

	for i := len(deckValues) - 1; i > 0; i-- {
		for j, suit := range deckValues[i] {
			if suit == 1 {
				straightFlushBuilder[j] = append(straightFlushBuilder[j], deck.Card{Val: i, Suit: j})
				flushBuilder[j] = append(flushBuilder[j], deck.Card{Val: i, Suit: j})
				suitIndexes = append(suitIndexes, j)
			} else {
				straightFlushBuilder[j] = []deck.Card{}
			}
		}

		// Straight flush
		for j := range straightFlushBuilder {
			if len(straightFlushBuilder[j]) == 5 {
				combinations["straight flush"] = straightFlushBuilder[j]
				out = true
				break
			}
		}

		if out {
			break
		}

		// Straight
		if len(suitIndexes) > 0 {
			straightBuilder = append(straightBuilder, deck.Card{Val: i, Suit: suitIndexes[0]})

			if len(straightBuilder) == 5 && len(combinations["straight"]) == 0 {
				combinations["straight"] = straightBuilder
			}
		} else {
			straightBuilder = []deck.Card{}
		}

		// Flush
		for j := range flushBuilder {
			if len(flushBuilder[j]) == 5 && len(combinations["flush"]) == 0 {
				combinations["flush"] = flushBuilder[j]
			}
		}

		// Pairs
		if len(suitIndexes) == 2 {
			key = "pair"
		} else if len(suitIndexes) == 3 {
			key = "three"
		} else if len(suitIndexes) == 4 {
			key = "four"
		}

		//fmt.Printf("%v\n", combinations)

		if key == "pair" && len(combinations["pair"]) != 0 {
			combinations["two pairs"] = combinations["pair"]
			//combinations["two pairs"] = append(combinations["two pairs"], combinations["pair"]...)
			key = "two pairs"
		} else if (key == "pair" && len(combinations["three"]) != 0) || (key == "three" && len(combinations["pair"]) != 0) {
			combinations["full house"] = combinations["three"]
			key = "full house"
		}

		if (len(combinations[key]) == 0 || key == "two pairs" || key == "full house") && key != "" {
			for _, suitIndex := range suitIndexes {
				combinations[key] = append(combinations[key], deck.Card{Val: i, Suit: suitIndex})
			}
		}

		suitIndexes = []int{}
		key = ""
	}

	/* for name, comb := range combinations {
		fmt.Printf("%v:%v\n", name, comb)
	} */

	resultKey, resultComb := "", []deck.Card{}

	for _, comb := range combinationsPriority {
		if len(combinations[comb]) > 0 {
			resultKey = comb
			resultComb = combinations[comb]
			break
		}
	}

	sort.SliceStable(resultComb, func(i, j int) bool {
		if resultComb[i].Val != resultComb[j].Val {
			return resultComb[i].Val < resultComb[j].Val
		}

		return resultComb[i].Suit < resultComb[j].Suit
	})
	//fmt.Printf("%v:%v\n", resultKey, resultComb)

	if resultKey == "straight flush" && resultComb[4].Val == 13 {
		resultKey = "royal flush"
	}

	return resultKey, resultComb
}

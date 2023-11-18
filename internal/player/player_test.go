package player

import (
	"fmt"
	"testing"

	"github.com/Scullder/poker/internal/deck"
)

/*func TestEvalComboRoyalFlush(t *testing.T) {

}

func TestEvalComboStraightFlush(t *testing.T) {

}

func TestEvalComboFour(t *testing.T) {

}

func TestEvalComboFullHouse(t *testing.T) {

}

func TestEvalComboFlush(t *testing.T) {

}

func TestEvalComboStraight(t *testing.T) {

}

func TestEvalComboThree(t *testing.T) {

}

func TestEvalComboTwoPair(t *testing.T) {

}

func TestEvalComboPair(t *testing.T) {

} */

/*
	0: "♥", // чирвы
	1: "♠", // пики
	2: "♦", // бубны
	3: "♣", // трефы
*/

var combinations = map[string][]deck.Card{
	"royal flush": {
		{Suit: 1, Val: 1},
		{Suit: 1, Val: 11},
		{Suit: 1, Val: 13},
		{Suit: 1, Val: 12},
		{Suit: 2, Val: 12},
		{Suit: 1, Val: 10},
		{Suit: 1, Val: 9},
	},
	"straight flush": {
		{Suit: 1, Val: 1},
		{Suit: 1, Val: 2},
		{Suit: 1, Val: 3},
		{Suit: 1, Val: 4},
		{Suit: 1, Val: 5},
		{Suit: 1, Val: 6},
		{Suit: 1, Val: 7},
	},
	"four": {
		{Suit: 3, Val: 1},
		{Suit: 1, Val: 2},
		{Suit: 1, Val: 7},
		{Suit: 0, Val: 7},
		{Suit: 2, Val: 7},
		{Suit: 1, Val: 6},
		{Suit: 3, Val: 7},
	},
	"full house": {
		{Suit: 1, Val: 1},
		{Suit: 1, Val: 13},
		{Suit: 2, Val: 13},
		{Suit: 1, Val: 2},
		{Suit: 3, Val: 13},
		{Suit: 2, Val: 2},
		{Suit: 1, Val: 7},
	},
	"flush": {
		{Suit: 1, Val: 12},
		{Suit: 2, Val: 2},
		{Suit: 1, Val: 3},
		{Suit: 0, Val: 4},
		{Suit: 1, Val: 5},
		{Suit: 1, Val: 2},
		{Suit: 1, Val: 7},
	},
	"straight": {
		{Suit: 2, Val: 6},
		{Suit: 3, Val: 8},
		{Suit: 1, Val: 10},
		{Suit: 2, Val: 7},
		{Suit: 0, Val: 5},
		{Suit: 0, Val: 6},
		{Suit: 3, Val: 9},
	},
	"three": {
		{Suit: 3, Val: 12},
		{Suit: 1, Val: 2},
		{Suit: 2, Val: 12},
		{Suit: 0, Val: 12},
		{Suit: 2, Val: 5},
		{Suit: 2, Val: 1},
		{Suit: 1, Val: 7},
	},
	"two pairs": {
		{Suit: 2, Val: 1},
		{Suit: 3, Val: 7},
		{Suit: 0, Val: 3},
		{Suit: 1, Val: 4},
		{Suit: 0, Val: 4},
		{Suit: 2, Val: 7},
		{Suit: 3, Val: 9},
	},
	"pair": {
		{Suit: 2, Val: 2},
		{Suit: 3, Val: 2},
		{Suit: 0, Val: 3},
		{Suit: 1, Val: 6},
		{Suit: 0, Val: 4},
		{Suit: 2, Val: 7},
		{Suit: 3, Val: 9},
	},
}

var expected = map[string][]deck.Card{
	"royal flush": {
		{Suit: 1, Val: 1},
		{Suit: 1, Val: 10},
		{Suit: 1, Val: 11},
		{Suit: 1, Val: 12},
		{Suit: 1, Val: 13},
	},
	"straight flush": {
		{Suit: 1, Val: 3},
		{Suit: 1, Val: 4},
		{Suit: 1, Val: 5},
		{Suit: 1, Val: 6},
		{Suit: 1, Val: 7},
	},
	"four": {
		{Suit: 0, Val: 7},
		{Suit: 1, Val: 7},
		{Suit: 2, Val: 7},
		{Suit: 3, Val: 7},
	},
	"full house": {
		{Suit: 1, Val: 2},
		{Suit: 2, Val: 2},
		{Suit: 1, Val: 13},
		{Suit: 2, Val: 13},
		{Suit: 3, Val: 13},
	},
	"flush": {
		{Suit: 1, Val: 2},
		{Suit: 1, Val: 3},
		{Suit: 1, Val: 5},
		{Suit: 1, Val: 7},
		{Suit: 1, Val: 12},
	},
	"straight": {
		{Suit: 0, Val: 6},
		{Suit: 2, Val: 7},
		{Suit: 3, Val: 8},
		{Suit: 3, Val: 9},
		{Suit: 1, Val: 10},
	},
	"three": {
		{Suit: 0, Val: 12},
		{Suit: 2, Val: 12},
		{Suit: 3, Val: 12},
	},
	"two pairs": {
		{Suit: 0, Val: 4},
		{Suit: 1, Val: 4},
		{Suit: 2, Val: 7},
		{Suit: 3, Val: 7},
	},
	"pair": {
		{Suit: 2, Val: 2},
		{Suit: 3, Val: 2},
	},
}

func TestEvalCombo(t *testing.T) {
	for key, cards := range combinations {
		actualName, actualCards := EvalCombo(cards)

		if actualName != key {
			t.Errorf("Combination name \"%v\" not equal to expected \"%v\"", actualName, key)
		}

		strActualCards := fmt.Sprintf("%v", actualCards)
		strExpectedCards := fmt.Sprintf("%v", expected[key])

		if strActualCards != strExpectedCards {
			t.Errorf("Combination %v: %v not equal to expected %v", key, strActualCards, strExpectedCards)
		}
	}
}

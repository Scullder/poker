package main

import (
	"github.com/Scullder/poker/internal/deck"
	"github.com/Scullder/poker/internal/player"
)

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

func main() {
	player.EvalCombo(combinations["royal flush"])
}

package main

import (
	"fmt"

	"github.com/Scullder/poker/internal/deck"
	"github.com/Scullder/poker/internal/player"
)

func main() {
	//deck := deck.MakeDeck()
	//fmt.Println(deck.Cards)

	player := player.MakePlayer("Boba")

	player.SetHand([]deck.Card{
		{Suit: 1, Val: 2},
		{Suit: 3, Val: 2},
	})

	cards := []deck.Card{
		{Suit: 1, Val: 1},
		{Suit: 2, Val: 2},
		{Suit: 1, Val: 3},
		{Suit: 1, Val: 4},
	}

	rank, err := player.EvalCombo(cards)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Hand rank is %v\n", rank)
}

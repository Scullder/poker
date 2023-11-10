package deck

import "fmt"

type Card struct {
	Suit int
	Val  int
}

func (c Card) String() string {
	return fmt.Sprintf("%v-%v", SuiteSymb(c.Suit), c.Val)
}

func SuiteSymb(num int) string {
	suites := [4]string{
		"♥", // чирвы
		"♠", // пики
		"♦", // бубны
		"♣", // трефы
	}

	return suites[num]
}

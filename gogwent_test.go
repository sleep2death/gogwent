package gogwent

import (
	"fmt"
	"testing"
)

func TestCards(t *testing.T) {
	cards, err := CreateCards("./assets/data/cards.json")
	if err != nil {
		t.Failed()
	}

	fmt.Printf("There are: '%d' cards in the deck.\n", len(*cards))
}

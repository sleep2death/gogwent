package gogwent

import (
	"fmt"
	"github.com/bmizerany/assert"
	"math/rand"
	"testing"
	"time"
)

func TestDeck(t *testing.T) {
	deck, err := NewDeckFromJSON("./assets/data/cards.json")

	assert.Equal(t, nil, err)
	assert.Equal(t, 300, cap(deck))
	assert.Equal(t, 36, deck.GetCardsNumber())
	assert.Equal(t, "Poor Fucking Infantry", deck.GetCardByID(1).name)
	assert.Equal(t, 6, deck.GetCardByName("Kaedweni Siege Expert").id)

	ok := deck.AddCard(&Card{id: 301, name: "Aspirin Shi"})
	if ok {
		card := deck.GetCardByID(301) //it's a pointer
		card.name = "Hello, Aspirin"
	}
	assert.Equal(t, "Hello, Aspirin", deck.GetCardByID(301).name)
	assert.Equal(t, 37, len(deck))

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	deck.Shuffle(r)

	fmt.Println(deck)
}

//Package gogwent implements a go library for cards game "Gwent"
package gogwent

import (
	"fmt"
	"github.com/bitly/go-simplejson" //handle json pase
	"io/ioutil"
	"math/rand"
	"strconv"
)

const (
	northern, nilfgaard, scoiatael, monsters, neutral = 1, 2, 3, 4, 0
	maxCardsInDeck                                    = 300
)

/*
Card holds the original data of a card
*/
type Card struct {
	id    int
	name  string
	img   string
	deck  uint
	power uint
	field uint
}

func jsonToCard(js *simplejson.Json) (*Card, bool) {
	id, err := strconv.ParseInt(js.Get("Id").MustString("-1"), 10, 0)
	if err != nil {
		return nil, false
	}

	name := js.Get("Name").MustString("")
	if name == "" {
		return nil, false
	}

	return &Card{id: int(id), name: name}, true
}

func (card Card) String() string {
	return fmt.Sprintf("Card[id:%v name:%v]", card.id, card.name)
}

//Deck holds a list of the cards
type Deck []*Card

//Shuffle cards in the deck
func (deck *Deck) Shuffle(r *rand.Rand) {
	data := *deck
	l := len(data)
	for i := 0; i < l; i++ {
		j := r.Int() % (i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

//GetCardsNumber return the number of all cards in the deck
func (deck *Deck) GetCardsNumber() int {
	return len(*deck)
}

//GetCardByID find the card by the given id
func (deck *Deck) GetCardByID(id int) *Card {
	for _, v := range *deck {
		if v.id == id {
			return v
		}
	}

	return nil
}

//GetCardByName find the first card which is matched the given name
func (deck *Deck) GetCardByName(name string) *Card {
	for _, v := range *deck {
		if v.name == name {
			return v
		}
	}

	return nil
}

//AddCard to the deck
func (deck *Deck) AddCard(card *Card) (ok bool) {
	ok = false
	if deck.GetCardByID(card.id) == nil && len(*deck) < maxCardsInDeck {
		*deck = append(*deck, card)
		ok = true
	}

	return
}

//NewDeckFromJSON creates all cards available in the game from an external json file
func NewDeckFromJSON(path string) (Deck, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	js, err := simplejson.NewJson(bytes)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	arr := js.MustArray()
	deck := make(Deck, 0, maxCardsInDeck)

	for i := range arr {
		card, ok := jsonToCard(js.GetIndex(i))
		if ok {
			deck.AddCard(card)
		}
	}

	return deck, nil
}

//NewDeck creates an empty deck with certain cap
func NewDeck() Deck {
	return make(Deck, 0, maxCardsInDeck)
}

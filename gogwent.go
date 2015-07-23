//Package gogwent implements a go library for cards game "Gwent"
package gogwent

import (
	"fmt"
	"github.com/bitly/go-simplejson" //handle json pase
	"io/ioutil"
	"strconv"
)

const (
	northern, nilfgaard, scoiatael, monsters, neutral = 1, 2, 3, 4, 0
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

//Deck holds a list of some cards
type Deck *[]*Card

//CreateCards creates all cards available in the game from an external file
func CreateCards(path string) (Deck, error) {
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
	list := make([]*Card, len(arr))

	for i := range arr {
		card, ok := jsonToCard(js.GetIndex(i))
		if ok {
			list[i] = card
		}
	}

	return &list, nil
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
	return fmt.Sprintf("Card[%v %v;]", card.id, card.name)
}

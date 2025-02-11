package deckofcards

import (
	"math/rand"
	"time"
)

// globals for suits and values
// go lets us declare these even if we don't use them jokerface.png
var suits []string = []string{"Clubs", "Diamonds", "Hearts", "Spades"}
var values []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

// Card represents a single playing card
type Card struct {
	Suit  string
	Value string
}

// CardDeck represents a deck of playing cards
type CardDeck struct {
	Cards []Card
}

// NewDeck initializes a new deck of cards in standard order
func NewDeck() *CardDeck {

	var cards []Card

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {

			card := Card{
				Suit:  suits[i],
				Value: values[j],
			}

			cards = append(cards, card)
		}
	}

	return &CardDeck{Cards: cards}
}

// Shuffle randomizes the order of the cards in the deck
func (d *CardDeck) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	for i := len(d.Cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)

		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}

}

// Contains checks if the deck contains a specific card
func (d *CardDeck) Contains(card Card) bool {

	for i := 0; i < len(d.Cards); i++ {
		if d.Cards[i] == card {
			return true
		}
	}

	return false
}

// DrawTop removes and returns the top card from the deck
func (d *CardDeck) DrawTop() Card {

	if len(d.Cards) == 0 {
		return Card{}
	}
	topCard := d.Cards[0]

	d.Cards = d.Cards[1:]

	return topCard
}

// DrawBottom removes and returns the bottom card from the deck
func (d *CardDeck) DrawBottom() Card {

	if len(d.Cards) == 0 {
		return Card{}
	}

	bottomCard := d.Cards[len(d.Cards)-1]

	d.Cards = d.Cards[:len(d.Cards)-1]

	return bottomCard
}

// DrawRandom removes and returns a card from a random position in the deck
func (d *CardDeck) DrawRandom() Card {

	if len(d.Cards) == 0 {
		return Card{}
	}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(d.Cards))

	randomCard := d.Cards[randomIndex]

	d.Cards = append(d.Cards[:randomIndex], d.Cards[randomIndex+1:]...)

	return randomCard
}

// CardToTop places a card on top of the deck
func (d *CardDeck) CardToTop(card Card) {

	d.Cards = append([]Card{card}, d.Cards...)

}

// CardToBottom places a card on the bottom of the deck
func (d *CardDeck) CardToBottom(card Card) {

	d.Cards = append(d.Cards, []Card{card}...)

}

// CardToRandom places a card at a random position in the deck
func (d *CardDeck) CardToRandom(card Card) {

	rand.Seed(time.Now().UnixNano())

	var random int = rand.Intn(len(d.Cards) + 1)

	d.Cards = append(d.Cards[:random], append([]Card{card}, d.Cards[random:]...)...)

}

// CardsLeft returns the number of cards left in the deck
func (d *CardDeck) CardsLeft() int {

	return len(d.Cards)
}

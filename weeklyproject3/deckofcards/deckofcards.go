package deckofcards

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

	return &CardDeck{}
}

// Shuffle randomizes the order of the cards in the deck
func (d *CardDeck) Shuffle() {

}

// Contains checks if the deck contains a specific card
func (d *CardDeck) Contains(card Card) bool {

	return false
}

// DrawTop removes and returns the top card from the deck
func (d *CardDeck) DrawTop() Card {
	return Card{}
}

// DrawBottom removes and returns the bottom card from the deck
func (d *CardDeck) DrawBottom() Card {
	return Card{}
}

// DrawRandom removes and returns a card from a random position in the deck
func (d *CardDeck) DrawRandom() Card {
	return Card{}
}

// CardToTop places a card on top of the deck
func (d *CardDeck) CardToTop(card Card) {

}

// CardToBottom places a card on the bottom of the deck
func (d *CardDeck) CardToBottom(card Card) {

}

// CardToRandom places a card at a random position in the deck
func (d *CardDeck) CardToRandom(card Card) {

}

// CardsLeft returns the number of cards left in the deck
func (d *CardDeck) CardsLeft() int {
	return 0
}

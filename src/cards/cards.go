package cards

import (
	"math/rand"
	"fmt"
	"sync"
)

//go:generate stringer -type=Suit
type Suit int

//go:generate stringer -type=Rank
type Rank int

const (
	Hearts   Suit = iota
	Spades
	Diamonds
	Clubs
)

const (
	Ace   Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func RankFromString(s string) Rank {
	switch s {
	case "Ace", "A":
		return Ace
	case "Two","2":
		return Two
	case "Three","3":
		return Three
	case "Four","4":
		return Four
	case "Five","5":
		return Five
	case "Six","6":
		return Six
	case "Seven","7":
		return Seven
	case "Eight","8":
		return Eight
	case "Nine","9":
		return Nine
	case "Ten","10":
		return Ten
	case "Jack","J":
		return Jack
	case "Queen","Q":
		return Queen
	case "King","K":
		return King
	}

	return -1
}

type Card struct {
	Rank Rank
	Suit Suit
}

type Deck struct {
	name  string
	Cards []Card
	mux   sync.Mutex
}

func NewHand(name string) *Hand {
	hand := new(Hand)
	hand.Cards = nil
	hand.name = name
	return hand
}

func NewDeck(name string) *Deck {
	deck := new(Deck)
	deck.initDeck()
	deck.name = name
	return deck
}

func (this *Deck) Take() Card {
	card := this.removeTopCard()
	return card
}

func (this *Deck) Add(card Card) {
	this.Cards = append(this.Cards, card)
}

func (this *Deck) Size() int {
	return len(this.Cards)
}

func (this *Deck) RandomCard() Card {
	return this.Cards[rand.Intn(len(this.Cards))]
}

func (this *Deck) RemoveAllCardsWithValue(rank Rank) {
	count, l := 0, len(this.Cards)
	for i := 0; i < l; i++ {
		c := this.Cards[i]
		if c.Rank == rank {
			this.removeCardAtIndex(i)
			count++
		}
		l = len(this.Cards)
	}
	this.print("Removed ", count, "cards with rank", rank)

}

func (this *Deck) removeCardAtIndex(index int) Card {
	this.mux.Lock()
	card := this.Cards[index]
	this.Cards = append(this.Cards[:index], this.Cards[index+1:]...)
	this.mux.Unlock()
	return card
}

func (this *Deck) removeRandomCard() Card {
	index := rand.Intn(len(this.Cards) - 1)
	card := this.Cards[index]
	this.removeCardAtIndex(index)

	return card
}

func (this *Deck) print(s ...interface{}) {
	fmt.Println(this.name, s)
}

func (this *Deck) Shuffle() {
	cards := this.Cards
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}
func (this *Deck) removeTopCard() Card {
	return this.removeCardAtIndex(0)
}

type Hand struct {
	Deck
}

func (this *Hand) HasCardWithValue(rank Rank) bool {
	for _, c := range this.Cards {
		if c.Rank == rank {
			return true
		}
	}
	return false
}

func (this *Hand) AddCard(card Card) {
	this.Add(card)
}

func (this *Deck) initDeck() {
	this.Cards = nil

	for suit := Hearts; suit <= Clubs; suit++ {
		for rank := Ace; rank <= King; rank++ {
			card := Card{Rank: rank, Suit: suit}
			this.Cards = append(this.Cards, card)
		}
	}
}

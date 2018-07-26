package cards

import (
	"math/rand"
	"fmt"
	"sync"
)

type Suit int

const (
	Hearts   Suit = iota
	Spades
	Diamonds
	Clubs
)

type Rank struct {
	value uint8
	name  string
	short string
}

var Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King =
	Rank{1, "Ace", "A"},
	Rank{2, "Two", "2"},
	Rank{3, "Three", "3"},
	Rank{4, "Four", "4"},
	Rank{5, "Five", "5"},
	Rank{6, "Six", "6"},
	Rank{7, "Seven", "7"},
	Rank{8, "Eight", "8"},
	Rank{9, "Nine", "9"},
	Rank{10, "Ten", "10"},
	Rank{11, "Jack", "J"},
	Rank{12, "Queen", "Q"},
	Rank{13, "King", "K"}

var Ranks = [13]Rank {
	Ace,Two,Three,Four,Five,Six,Seven,Eight,Nine,Ten,Jack,Queen,King }

func FromString(s string) Rank {

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

	for _, suit := range Suits {
		for _, rank := range Ranks {
			card := Card{Rank: rank, Suit: suit}
			this.Cards = append(this.Cards, card)
		}
	}
}

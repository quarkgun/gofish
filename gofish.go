package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

var cardSuites = [4]string{"H", "S", "D", "C"}
var cardValues = [13]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "J", "Q", "K", "A"}

var playerHand, cpuHand Hand
var playerScore, cpuScore int
var deck Deck

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Go Fish!")

	deck = *NewDeck()
	playerHand = *new(Hand)
	cpuHand = *new(Hand)

	db("Dealing...")
	deal()

	playerTurn := true
	quit := false
	var input string

	for !quit {
		db("Your Hand", playerHand)
		if playerTurn {
			fmt.Print("Choose a card: ")
			fmt.Scanln(&input)

			if strings.EqualFold(input, "quit") {
				quit = true
				break
			} else {
				db("Do you have a %v?", input)
				time.Sleep(time.Second * 1)
				if cpuHand.hasCard(input) {
					db("Yep")

				} else {
					db("Go Fish!")
					playerTurn = false
				}
			}


		} else {
			db("CPU turn...")
			val := cpuHand.randomCard().value
			db("Do you have a %v?", val)
			time.Sleep(time.Second * 1)

		}
	}

}

func db(f string, s ...interface{}) {
	fmt.Printf(f, s)
}

func deal() {
	for i := 0; i < 5; i++ {
		playerHand.addCard(deck.take())
		cpuHand.addCard(deck.take())
	}
}

func (this *Deck) initDeck() {
	this.cards = make(map[string][]string)

	for _, s := range cardSuites {
		for _, val := range cardValues {
			this.cards[val] = append(this.cards[val], s)
		}
	}
}

type Card struct {
	value string
	suite string
}

type Deck struct {
	cards map[string][]string
}

func NewDeck() *Deck {
	deck := new(Deck)
	deck.initDeck()
	return deck
}

func (this *Deck) take() Card {
	defer func() {
		this.cards = this.cards[1:]
	}()
	return this.cards[0]
}

func (this *Deck) add(card Card) {
	this.cards = append(this.cards, card)
}

func (this *Deck) peek() Card {
	return this.cards[0]
}

func (this *Deck) size() int {
	count := 0
	for _, num := range this.cards {
		count += len(num)
	}
	return count
}

func (this *Deck) randomCard() Card {
//	return this.cards[rand.Intn(len(this.cards))]

	


}

func (this *Deck) removeAllCardsWithValue(value string) {
	for _, c := range this.cards {

	}
}

type Hand struct {
	Deck
}

func (this *Hand) hasCard(val string) bool {
	for _, c := range this.cards {
		if c.value == val {
			return true
		}
	}
	return false
}

func (this *Hand) addCard(card Card) {
	this.add(card)
}

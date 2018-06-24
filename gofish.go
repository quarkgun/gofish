package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
	"sync"
)

var SUITES = [4]string{"Hearts", "Spades", "Diamonds", "Clubs"}
var VALUES = [13]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "J", "Q", "K", "A"}

var playerHand, cpuHand Hand
var playerScore, cpuScore = 0, 0
var deck Deck

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Go Fish!")

	deck = *NewDeck("deck")
	playerHand = *NewHand("playerHand")
	cpuHand = *NewHand("cpuHand")

	deck.shuffle()

	fmt.Println("Dealing...")
	deal()

	playerTurn := true
	quit := false
	var input string

	for !quit {
		checkDeck()
		fmt.Printf("Scores: Player[%v] CPU[%v]\n", playerScore, cpuScore)
		fmt.Printf("Your Hand: %v\n", playerHand.size())
		fmt.Println(playerHand)
		if playerTurn {
			fmt.Println("Your turn!")
			time.Sleep(time.Second * 1)
			fmt.Print("\n\nChoose a card: ")
			fmt.Scanln(&input)

			if strings.EqualFold(input, "quit") {
				quit = true
				break
			} else {
				fmt.Printf("Do you have a %v?\n", input)
				time.Sleep(time.Second * 1)
				if cpuHand.hasCard(input) {
					fmt.Println("Yep!")
					time.Sleep(time.Second * 1)
					playerHand.removeAllCardsWithValue(input)
					cpuHand.removeAllCardsWithValue(input)
					playerScore++

				} else {
					fmt.Println("Go Fish!")
					time.Sleep(time.Second * 1)
					playerTurn = false
					playerHand.add(deck.take())
				}
			}

		} else {
			time.Sleep(time.Second * 1)
			fmt.Println("\n\nCPU turn:")

			val := cpuHand.randomCard().value
			fmt.Printf("Do you have a %v?\n", val)

			if playerHand.hasCard(val) {
				fmt.Println("Yep!")
				time.Sleep(time.Second * 1)
				playerHand.removeAllCardsWithValue(val)
				cpuHand.removeAllCardsWithValue(val)
				cpuScore++
			} else {
				fmt.Println("Go Fish!")
				time.Sleep(time.Second * 1)
				playerTurn = true
				cpuHand.add(deck.take())
			}
		}
	}
}

func db(v ...interface{}) {
	fmt.Println(v...)
}

func dbv(f string, v ...interface{}) {
	fmt.Printf(f, v...)
}

func checkDeck() {
	for i, c := range deck.cards {
		if c.value == "" {
			dbv("ERROR: deck(%[1]d) has empty card at index %[2]d\n", len(deck.cards), i)
		}
	}
}

func deal() {
	for i := 0; i < 5; i++ {
		playerHand.addCard(deck.take())
		cpuHand.addCard(deck.take())
	}
}

func (this *Deck) initDeck() {
	this.cards = nil

	for _, s := range SUITES {
		for _, val := range VALUES {
			card := Card{val, s}
			this.cards = append(this.cards, card)
		}
	}
}

type Card struct {
	value string
	suite string
}

type Deck struct {
	name  string
	cards []Card
	mux   sync.Mutex
}

func NewHand(name string) *Hand {
	hand := new(Hand)
	hand.cards = nil
	hand.name = name
	return hand
}

func NewDeck(name string) *Deck {
	deck := new(Deck)
	deck.initDeck()
	deck.name = name
	return deck
}

func (this *Deck) take() Card {
	card := this.removeTopCard()
	return card
}

func (this *Deck) add(card Card) {
	this.cards = append(this.cards, card)
}

func (this *Deck) size() int {
	return len(this.cards)
}

func (this *Deck) randomCard() Card {
	return this.cards[rand.Intn(len(this.cards))]
}

func (this *Deck) removeAllCardsWithValue(value string) {
	count, l := 0, len(this.cards)
	for i := 0; i < l; i++ {
		c := this.cards[i]
		if c.value == value {
			this.removeCardAtIndex(i)
			count++
		}
		l = len(this.cards)
	}
	this.print("Removed ", count, "cards with value", value)

}

func (this *Deck) removeCardAtIndex(index int) Card {
	this.mux.Lock()
	card := this.cards[index]
	this.cards = append(this.cards[:index], this.cards[index+1:]...)
	this.mux.Unlock()
	return card
}

func (this *Deck) removeRandomCard() Card {
	index := rand.Intn(len(this.cards) - 1)
	card := this.cards[index]
	this.removeCardAtIndex(index)

	return card
}

func (this *Deck) print(s ...interface{}) {
	fmt.Println(this.name, s)
}

func (this *Deck) shuffle() {
	cards := this.cards
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
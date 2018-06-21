package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
	"sync"
	"os"
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
	fmt.Println("Shuffled deck:", deck)

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
			fmt.Println("Your turn!\n\n")
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
					playerTurn = false
					playerHand.add(deck.take())
				}
			}

		} else {
			fmt.Println("\n\nCPU turn...\n\n")
			time.Sleep(time.Second * 1)

			val := cpuHand.randomCard().value
			fmt.Printf("Do you have a %v?\n", val)
			time.Sleep(time.Second * 1)

			if playerHand.hasCard(val) {
				fmt.Println("Yep!")
				playerHand.removeAllCardsWithValue(val)
				cpuHand.removeAllCardsWithValue(val)
				cpuScore++
			} else {
				fmt.Println("Go Fish!")
				playerTurn = true
				cpuHand.add(deck.take())
			}
			time.Sleep(time.Second * 2)
		}
	}
}

func db(v ...interface{}) {
	fmt.Println(v)
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

func test() {
	deck := *NewDeck("testDeck")
	fmt.Println(deck)
	deck.removeCardAtIndex(2)
	fmt.Println(deck)
	os.Exit(0)
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
			fmt.Println("Created card", card)
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
	card := this.removeRandomCard()
	return card
}

func (this *Deck) add(card Card) {
	this.print("Adding card", card)
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
	this.print("Removing card at index", index, card)
	if index == 0 {
		this.cards = this.cards[1:]
	} else {
		tmp1 := make([]Card, index-1)
		tmp2 := make([]Card, len(this.cards)-index)

		copy(tmp1, this.cards[:index])
		copy(tmp2, this.cards[index+1:])
		this.cards = make([]Card, len(tmp1) + len(tmp2))
		this.cards = append(tmp1, tmp2...)
	}
	this.mux.Unlock()
	return card
}

func (this *Deck) removeRandomCard() Card {
	index := rand.Intn(len(this.cards)-1)
	card := this.cards[index]
	this.print("Removed random card", card)
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

func IndexOf(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Contains(vs []string, t string) bool {
	return IndexOf(vs, t) >= 0
}

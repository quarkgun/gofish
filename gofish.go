package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

var SUITES = [4]string{"Hearts", "Spades", "Diamonds", "Clubs"}
var VALUES = [13]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "J", "Q", "K", "A"}

var playerHand, cpuHand Hand
var playerScore, cpuScore int = 0,0
var deck Deck

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Go Fish!")

	deck = *NewDeck()
	playerHand = *NewHand()
	cpuHand = *NewHand()

	fmt.Println("Dealing...")
	deal()

	playerTurn := true
	quit := false
	var input string

	for !quit {
		fmt.Printf("Scores: Player[%v] CPU[%v]\n", playerScore, cpuScore)
		fmt.Printf("Your Hand: %v\n", playerHand.size())
		fmt.Println(playerHand)
		fmt.Println(cpuHand)
		if playerTurn {
			fmt.Print("Choose a card: ")
			fmt.Scanln(&input)

			if strings.EqualFold(input, "quit") {
				quit = true
				break
			} else {
				fmt.Printf("Do you have a %v?", input)
				time.Sleep(time.Second * 1)
				if cpuHand.hasCard(input) {
					fmt.Println("Yep")
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
			fmt.Println("CPU turn...")
			val := cpuHand.randomCard().value
			fmt.Printf("Do you have a %v?", val)
			time.Sleep(time.Second * 1)

			if playerHand.hasCard(val) {
				fmt.Println("Yep")
				playerHand.removeAllCardsWithValue(input)
				cpuHand.removeAllCardsWithValue(input)
				cpuScore++
			} else {
				fmt.Println("Go Fish!")
				playerTurn = true
				cpuHand.add(deck.take())
			}
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
	this.cards = make(map[string][]string)

	for _, s := range SUITES {
		for _, val := range VALUES {
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

func NewHand() *Hand {
	hand := new(Hand)
	hand.cards = make(map[string][]string)
	return hand
}

func NewDeck() *Deck {
	deck := new(Deck)
	deck.initDeck()
	return deck
}

func (this *Deck) take() Card {
	card := this.randomCard()
	this.removeCard(card)
	fmt.Println("Taking card", card)
	return card
}

func (this *Deck) add(card Card) {
	this.cards[card.value] = append(this.cards[card.value], card.suite)
}

func (this *Deck) size() int {
	return len(this.cards)
}

func (this *Deck) randomCard() Card {
//	return this.cards[rand.Intn(len(this.cards))]
	value := randMapKey(this.cards)
	suites := this.cards[value]

	if len(suites) > 0 {
		suite := randListVal(suites)
		return Card{value:value, suite:suite}
	}
	panic("Tried to get an invalid card")
}

func randListVal(list []string) string {
	return list[rand.Intn(len(list))]
}

func randMapKey(m map[string][]string) string {
	i := rand.Intn(len(m))
	for k := range m {
		if i == 0 {
			return k
		}
		i--
	}
	panic("never")
}

func (this *Deck) removeAllCardsWithValue(value string) {
	delete(this.cards, value)
}

func (this *Deck) removeCard(card Card) {
	values := this.cards[card.value]
	for i, s := range values {
		if s == card.suite {
			deleteElementAt(values, i)
		}
	}
}

func deleteElementAt(slice []string, index int) {
	slice[index] = slice[len(slice)-1] // Copy lslicest element to index i.
	slice[len(slice)-1] = ""   // Erslicese lslicest element (write zero vslicelue).
	slice = slice[:len(slice)-1]   // Truncslicete slice.
}

type Hand struct {
	Deck
}

func (this *Hand) hasCard(val string) bool {
	values := this.cards[val]
	return len(values) > 0
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
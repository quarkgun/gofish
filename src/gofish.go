package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
	. "cards"
)

var playerHand, cpuHand Hand
var playerScore, cpuScore = 0, 0
var deck Deck

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Go Fish!")

	deck = *NewDeck("deck")
	playerHand = *NewHand("playerHand")
	cpuHand = *NewHand("cpuHand")

	deck.Shuffle()

	fmt.Println("Dealing...")
	deal()

	playerTurn := true
	quit := false
	var input string

	for !quit {
		checkDeck()
		fmt.Printf("Scores: Player[%v] CPU[%v]\n", playerScore, cpuScore)
		fmt.Printf("Your Hand: %v\n", playerHand.Size())
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
				if cpuHand.HasCardWithValue(input) {
					fmt.Println("Yep!")
					time.Sleep(time.Second * 1)
					playerHand.RemoveAllCardsWithValue(input)
					cpuHand.RemoveAllCardsWithValue(input)
					playerScore++

				} else {
					fmt.Println("Go Fish!")
					time.Sleep(time.Second * 1)
					playerTurn = false
					playerHand.Add(deck.Take())
				}
			}

		} else {
			time.Sleep(time.Second * 1)
			fmt.Println("\n\nCPU turn:")

			val := cpuHand.RandomCard().Rank
			fmt.Printf("Do you have a %v?\n", val)

			if playerHand.HasCardWithValue(val) {
				fmt.Println("Yep!")
				time.Sleep(time.Second * 1)
				playerHand.RemoveAllCardsWithValue(val)
				cpuHand.RemoveAllCardsWithValue(val)
				cpuScore++
			} else {
				fmt.Println("Go Fish!")
				time.Sleep(time.Second * 1)
				playerTurn = true
				cpuHand.Add(deck.Take())
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
	for i, c := range deck.Cards {
		if c.Rank < Ace || c.Rank > King {
			dbv("ERROR: deck(%[1]d) has empty card at index %[2]d\n", len(deck.Cards), i)
		}
	}
}

func deal() {
	for i := 0; i < 5; i++ {
		playerHand.AddCard(deck.Take())
		cpuHand.AddCard(deck.Take())
	}
}


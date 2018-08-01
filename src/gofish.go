package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
	"cards"
)

var playerHand, cpuHand cards.Hand
var playerScore, cpuScore = 0, 0
var deck cards.Deck

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Go Fish!")

	deck = *cards.NewDeck("deck")
	playerHand = *cards.NewHand("playerHand")
	cpuHand = *cards.NewHand("cpuHand")

	deck.Shuffle()

	fmt.Println("Dealing...")
	deal()

	playerTurn := true
	quit := false
	var input string

	for !quit {
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

				rank := cards.RankFromString(input)
				if rank < cards.Ace || rank > cards.King {
					continue
				}

				fmt.Printf("Do you have a %v?\n", input)
				time.Sleep(time.Second * 1)

				if cpuHand.HasCardWithValue(rank) {
					fmt.Println("Yep!")
					time.Sleep(time.Second * 1)
					playerHand.RemoveAllCardsWithValue(rank)
					cpuHand.RemoveAllCardsWithValue(rank)
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


func deal() {
	for i := 0; i < 5; i++ {
		playerHand.AddCard(deck.Take())
		cpuHand.AddCard(deck.Take())
	}
}


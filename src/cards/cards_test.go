package cards

import (
	"testing"
	"fmt"
)

func TestRankFromString(t *testing.T) {
	var rank Rank
	rank = RankFromString("A")
	fmt.Printf("Got Rank %d", rank)
	if rank != Ace {
		t.Fail()
	}
}

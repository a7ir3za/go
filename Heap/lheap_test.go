package lheap

import (
	"log"
	"testing"
)

func Test2542(t *testing.T) {
	if MaxScore(
		[]int{1, 3, 3, 2},
		[]int{2, 1, 4, 3}, 3) != 12 {
		t.Fatalf("Wrong maxScore! must be: %d", 12)
	}
}

func Test948(t *testing.T) {
	if BagOfTokensScore([]int{100, 200, 300, 400}, 200) != 2 {
		log.Fatal("Wrong score! must be: 2")
	}

	for _, x := range [][][]int{{{200, 100}, {150, 1}}, {{100}, {50, 0}}, {{81, 91, 31}, {73, 1}}, {{25}, {150, 1}}} {
		log.Printf("+ Score: %d   Power: %d   Tokens: %v", x[1][1], x[1][0], x[0])
		if r := BagOfTokensScore(x[0], x[1][0]); r != x[1][1] {
			t.Errorf("Wrong score! must be: %d != %d", x[1][1], r)
		}
	}
}

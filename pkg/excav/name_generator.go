package excav

import (
	"math/rand"
	"time"
)

var (
	// nouns and adjectives are used for random
	// name generator
	adjectives = []string{"blurry", "red", "green", "fat", "slim", "tall", "slow", "fast", "hungry"}
	nouns      = []string{"fish", "lion", "cat", "dog", "squirrel", "sparrow", "elephant"}
)

func randomName() string {
	rand.Seed(time.Now().UnixNano())

	randAdj := adjectives[rand.Intn(len(adjectives))]
	randNoun := nouns[rand.Intn(len(nouns))]

	return randAdj + "-" + randNoun
}

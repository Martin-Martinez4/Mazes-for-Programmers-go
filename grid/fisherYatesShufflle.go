package grid

import "math/rand"

// Do the Fisher-Yates Shuffle
func shuffle[T any](s []T) {

	for i := 0; i < len(s); i++ {
		randIndex := rand.Intn(len(s))
		s[i], s[randIndex] = s[randIndex], s[i]
	}
}

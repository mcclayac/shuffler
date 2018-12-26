package shuffler

import (
	"math/rand"
)

var (
	counter int = 0
)

func init()  {
	counter = 0  // demonstration, it
				// is default initialized to zero

}

func GetCount() int  {
	return counter
}

type Shuffleable interface {
	Len() int
	Swap(i, j int)
}

type WeightedShuffleable interface {
	Shuffleable
	Weight(i int) int

}

func Shuffle(s Shuffleable) {
	for i:= 0; i < s.Len() ; i++ {
		j := rand.Intn(s.Len()-i)
		s.Swap(i, j)
		counter++
	}
}

func WeightedShuffle(w WeightedShuffleable) {
	total := 0
	for i:= 0; i < w.Len() ; i++ {
		total += w.Weight(i)
	}

	for i:= 0; i < w.Len() ; i++ {
		pos := rand.Intn(total)
		cum := 0
		for j := i; j < w.Len(); j++ {
			counter++
			cum += w.Weight(j)
			if pos >= cum {
				total -= w.Weight(j)
				w.Swap(i, j)
				break
			}
		}
	}
}

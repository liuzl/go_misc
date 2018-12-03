package main

import (
	"fmt"

	"github.com/dexyk/stringosim"
)

func main() {
	fmt.Println(stringosim.Levenshtein([]rune("stringosim"), []rune("stingobim")))

	fmt.Println(stringosim.Levenshtein([]rune("stringosim"), []rune("stingobim"),
		stringosim.LevenshteinSimilarityOptions{
			InsertCost:     3,
			DeleteCost:     5,
			SubstituteCost: 2,
		}))

	fmt.Println(stringosim.Levenshtein([]rune("stringosim"), []rune("STRINGOSIM"),
		stringosim.LevenshteinSimilarityOptions{
			InsertCost:      3,
			DeleteCost:      4,
			SubstituteCost:  5,
			CaseInsensitive: true,
		}))
}

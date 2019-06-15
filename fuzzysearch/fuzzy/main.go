package main

import (
	"fmt"
	"github.com/sajari/fuzzy"
)

func main() {
	model := fuzzy.NewModel()

	// For testing only, this is not advisable on production
	model.SetThreshold(1)

	// This expands the distance searched, but costs more resources (memory and time).
	// For spell checking, "2" is typically enough, for query suggestions this can be higher
	model.SetDepth(5)

	// Train multiple words simultaneously by passing an array of strings to the "Train" function
	words := []string{"bob", "your", "uncle", "dynamite", "delicate", "biggest", "big", "bigger", "aunty", "you're"}
	model.Train(words)

	// Train word by word (typically triggered in your application once a given word is popular enough)
	model.TrainWord("single")

	// Check Spelling
	fmt.Println("\nSPELL CHECKS")
	fmt.Println("	Deletion test (yor) : ", model.SpellCheck("yor"))
	fmt.Println("	Swap test (uncel) : ", model.SpellCheck("uncel"))
	fmt.Println("	Replace test (dynemite) : ", model.SpellCheck("dynemite"))
	fmt.Println("	Insert test (dellicate) : ", model.SpellCheck("dellicate"))
	fmt.Println("	Two char test (dellicade) : ", model.SpellCheck("dellicade"))

	// Suggest completions
	fmt.Println("\nQUERY SUGGESTIONS")
	fmt.Println("	\"bigge\". Did you mean?: ", model.Suggestions("bigge", false))
	fmt.Println("	\"bo\". Did you mean?: ", model.Suggestions("bo", false))
	fmt.Println("	\"dyn\". Did you mean?: ", model.Suggestions("dyn", false))

	// Autocomplete suggestions
	suggested, _ := model.Autocomplete("bi")
	fmt.Printf("	\"bi\". Suggestions: %v", suggested)

}

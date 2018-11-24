package main

import (
	"fmt"

	"github.com/pilosa/go-pilosa"
)

func main() {
	// We will just use the default client which assumes the server is at http://localhost:10101
	client := pilosa.DefaultClient()

	// Let's load the schema from the server.
	// Note that, for this example the schema should be created beforehand
	// and the stargazer data should be imported.
	// See the Getting Started repository: https://github.com/pilosa/getting-started/
	schema, err := client.Schema()
	if err != nil {
		// Most calls will return an error value.
		// You should handle them appropriately.
		// We will just terminate the program in this case.
		// Error handling was left out for brevity in the rest of the code.
		panic(err)
	}

	// We need to refer to indexes and fields before we can use them in a query.
	repository := schema.Index("repository")
	stargazer := repository.Field("stargazer")
	language := repository.Field("language")

	var response *pilosa.QueryResponse

	// Which repositories did user 14 star:
	response, _ = client.Query(stargazer.Row(14))
	fmt.Println("User 14 starred: ", response.Result().Row().Columns)

	// What are the top 5 languages in the sample data?
	response, err = client.Query(language.TopN(5))
	languageIDs := []uint64{}
	for _, item := range response.Result().CountItems() {
		languageIDs = append(languageIDs, item.ID)
	}
	fmt.Println("Top 5 languages: ", languageIDs)

	// Which repositories were starred by both user 14 and 19:
	response, _ = client.Query(
		repository.Intersect(
			stargazer.Row(14),
			stargazer.Row(19)))
	fmt.Println("Both user 14 and 19 starred:", response.Result().Row().Columns)

	// Which repositories were starred by user 14 or 19:
	response, _ = client.Query(
		repository.Union(
			stargazer.Row(14),
			stargazer.Row(19)))
	fmt.Println("User 14 or 19 starred:", response.Result().Row().Columns)

	// Which repositories were starred by user 14 or 19 and were written in language 1:
	response, _ = client.Query(
		repository.Intersect(
			repository.Union(
				stargazer.Row(14),
				stargazer.Row(19),
			),
			language.Row(1)))
	fmt.Println("User 14 or 19 starred, written in language 1:", response.Result().Row().Columns)

	// Set user 99999 as a stargazer for repository 77777?
	client.Query(stargazer.Set(99999, 77777))
}

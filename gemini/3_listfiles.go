package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	// Access your API key as an environment variable
	client, err := genai.NewClient(ctx,
		option.WithAPIKey(os.Getenv("API_KEY")),
		option.WithEndpoint(os.Getenv("ENDPOINT")),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	it := client.ListFiles(ctx)

	for {
		file, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Found file %s, %s, %s\n", file.URI, file.Name, file.DisplayName)
	}
}

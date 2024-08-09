package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/api/option"
)

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}

func main() {
	fmt.Println(os.Getenv("API_KEY"))
	fmt.Println(os.Getenv("ENDPOINT"))
	ctx := context.Background()
	client, err := genai.NewClient(ctx,
		option.WithAPIKey(os.Getenv("API_KEY")),
		option.WithEndpoint(os.Getenv("ENDPOINT")),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a AI and magic in Chinese."))
	if err != nil {
		log.Fatal(err)
	}
	printResponse(resp)
}

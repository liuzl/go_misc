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
	prompt := []genai.Part{
		genai.FileData{URI: "https://generativelanguage.googleapis.com/v1beta/files/suybti91qili"},
		//genai.Text("Can you summarize this document as a bulleted list? Please response in Chinese."),
		genai.Text("详细讲解一下最后一道题。"),
	}
	resp, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(resp)
}

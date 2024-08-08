package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	llm, err := openai.New(openai.WithBaseURL(os.Getenv("OPENAI_BASE_URL")), openai.WithToken(os.Getenv("OPENAI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llm.Call(ctx, "You should response in Chinese. The first man to walk on the moon",
		llms.WithTemperature(0.8),
		llms.WithStopWords([]string{"Armstrong"}),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)
}

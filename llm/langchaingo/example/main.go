package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {
	llm, err := openai.New(openai.WithBaseURL(os.Getenv("OPENAI_BASE_URL")), openai.WithToken(os.Getenv("OPENAI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	completion, err := llm.Call(ctx, "You should response in Chinese. 写一首关于北京炎炎夏日的诗。")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)
}

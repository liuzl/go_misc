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

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx,
		option.WithAPIKey(os.Getenv("API_KEY")),
		option.WithEndpoint(os.Getenv("ENDPOINT")),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Use client.UploadFile to upload a file to the service.
	// Pass it an io.Reader.
	f, err := os.Open("math.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Optionally set a display name.
	opts := genai.UploadFileOptions{DisplayName: "Math Exam PDF"}
	// Let the API generate a unique `name` for the file by passing an empty string.
	// If you specify a `name`, then it has to be globally unique.
	doc1, err := client.UploadFile(ctx, "", f, &opts)
	if err != nil {
		log.Fatal(err)
	}

	// View the response.
	fmt.Printf("Uploaded file %s as: %q\n", doc1.DisplayName, doc1.URI)
}

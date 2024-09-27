package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/vertexai/genai"
	_ "github.com/joho/godotenv/autoload"
)

func tryGemini(w io.Writer, projectID string, location string, modelName string) error {
	// location := "us-central1"
	// modelName := "gemini-1.0-pro-vision-001"

	ctx := context.Background()
	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		return fmt.Errorf("error creating client: %w", err)
	}
	gemini := client.GenerativeModel(modelName)

	img := genai.FileData{
		MIMEType: "image/jpeg",
		FileURI:  "gs://generativeai-downloads/images/scones.jpg",
	}
	prompt := genai.Text("What is in this image?")

	resp, err := gemini.GenerateContent(ctx, img, prompt)
	if err != nil {
		return fmt.Errorf("error generating content: %w", err)
	}
	rb, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return fmt.Errorf("json.MarshalIndent: %w", err)
	}
	fmt.Fprintln(w, string(rb))
	return nil
}

func main() {
	projectID := os.Getenv("PROJECT_ID")
	location := os.Getenv("LOCATION")
	modelName := os.Getenv("MODEL_NAME")

	err := tryGemini(os.Stdout, projectID, location, modelName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

//zliu-813@norse-carport-405405.iam.gserviceaccount.com

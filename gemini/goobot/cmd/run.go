/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a prompt through the Gemini AI model",
	Long: `This command allows you to run a prompt through the Gemini AI model.
It accepts one or two arguments:

1. With one argument: The argument is treated as the prompt text.
   Example: run "What is the capital of France?"

2. With two arguments: The first argument is a file path, and the second is the prompt text.
   The contents of the file will be used as context for the prompt.
   Example: run /path/to/file.txt "Summarize this document"

The command uses the Gemini 1.5 Flash model and supports code execution as a tool.
The response from the AI model will be printed to the console.`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
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
		model.Tools = []*genai.Tool{
			{CodeExecution: &genai.CodeExecution{}},
		}
		var prompt []genai.Part
		if len(args) == 1 {
			prompt = []genai.Part{genai.Text(args[0])}
		} else {
			prompt = []genai.Part{
				genai.FileData{URI: args[0]},
				genai.Text(args[1]),
			}
		}
		resp, err := model.GenerateContent(ctx, prompt...)
		if err != nil {
			log.Fatal(err)
		}
		printResponse(resp)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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

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
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List files associated with the Gemini AI model",
	Long: `This command lists all files associated with your Gemini AI model.

It displays information about each file, including:
- URI: The unique identifier for the file
- Name: The file name
- DisplayName: The human-readable name of the file

This command is useful for managing and reviewing the files you've uploaded
or associated with your Gemini AI model, which can be used as context
for prompts in other commands like 'run'.

Usage:
  gemini-cli list`,
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

		it := client.ListFiles(ctx)

		for {
			file, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("URI: %s, Name: %s, DisplayName: %s\n", file.URI, file.Name, file.DisplayName)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

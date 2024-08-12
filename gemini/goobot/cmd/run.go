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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

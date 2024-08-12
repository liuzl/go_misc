/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
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
		if err := client.DeleteFile(ctx, args[0]); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

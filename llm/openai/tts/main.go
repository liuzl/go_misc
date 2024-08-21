package main

import (
	"context"
	"fmt"
	"io"
	"os"

	_ "github.com/joho/godotenv/autoload"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	config := openai.DefaultConfig(os.Getenv("OPENAI_API_KEY"))
	config.BaseURL = os.Getenv("OPENAI_BASE_URL")
	client := openai.NewClientWithConfig(config)

	res, err := client.CreateSpeech(context.Background(), openai.CreateSpeechRequest{
		Model: openai.TTSModel1,
		Input: "你好，这是一个TTS的测试。我看打车到古北单程要300多。",
		Voice: "zliu", //openai.VoiceAlloy,
	})
	if err != nil {
		fmt.Printf("CreateSpeech error: %v\n", err)
		return
	}
	defer res.Close()
	buf, err := io.ReadAll(res)
	if err != nil {
		fmt.Printf("ReadAll error: %v\n", err)
		return
	}
	err = os.WriteFile("test.mp3", buf, 0644)
	if err != nil {
		fmt.Printf("WriteFile error: %v\n", err)
		return
	}
}

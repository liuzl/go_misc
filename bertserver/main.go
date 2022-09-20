package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/nlpodyssey/cybertron/pkg/tasks"
	"github.com/nlpodyssey/cybertron/pkg/tasks/text2text"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	modelsDir = "models"
	modelName = "Helsinki-NLP/opus-mt-en-zh"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	m, err := tasks.Load[text2text.Interface](&tasks.Config{ModelsDir: modelsDir, ModelName: modelName})
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer tasks.Finalize(m)

	opts := text2text.DefaultOptions()

	fn := func(text string) error {
		start := time.Now()
		result, err := m.Generate(context.Background(), text, opts)
		if err != nil {
			return err
		}
		fmt.Println(time.Since(start).Seconds())
		fmt.Println(result.Texts[0])
		return nil
	}

	err = ForEachInput(os.Stdin, fn)
	if err != nil {
		log.Fatal().Err(err)
	}
}

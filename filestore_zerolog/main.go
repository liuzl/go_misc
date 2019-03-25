package main

import (
	"flag"
	"os"
	"path/filepath"
	"sync"

	"github.com/liuzl/filestore"
	"github.com/rs/zerolog"
)

var (
	dir        = flag.String("dir", filepath.Join(filepath.Dir(os.Args[0]), "chargelog"), "charge log dir")
	chargeLog  *zerolog.Logger
	chargeOnce sync.Once
)

func getLogger() *zerolog.Logger {
	chargeOnce.Do(func() {
		f, err := filestore.NewFileStore(*dir)
		if err != nil {
			panic(err)
		}
		hostname, _ := os.Hostname()
		log := zerolog.New(f).With().
			Timestamp().
			Str("service", filepath.Base(os.Args[0])).
			Str("host", hostname).
			Str("method", "callback").
			Str("url", "/v1/iswhatsapp/callback").
			Logger()
		chargeLog = &log
	})
	return chargeLog
}

func main() {
	getLogger().Debug().Str("uname", "zhangzhong").Float64("Cost", 1).Msg("charge")
}

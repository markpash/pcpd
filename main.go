package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/markpash/pcpd/internal/server"
	"github.com/markpash/pcpd/internal/server/config"
)

func main() {
	cfgPath := flag.String("c", "pcpd.toml", "path to pcpd config file")
	flag.Parse()

	cfgFile, err := os.Open(*cfgPath)
	if err != nil {
		panic(err)
	}

	config, err := config.Parse(cfgFile)
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := server.Start(ctx, *config); err != nil {
		panic(err)
	}
}

func panic(err error) {
	fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}

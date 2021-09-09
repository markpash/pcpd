package main

import (
	"context"
	"fmt"
	"os"

	"github.com/markpash/pcpd/internal/server"
)

func main() {
	if err := server.Start(context.Background()); err != nil {
		panic(err)
	}
}

func panic(err error) {
	fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}

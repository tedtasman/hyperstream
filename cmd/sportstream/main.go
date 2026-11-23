package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/tedtasman/hyperstream/internal/control"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	path := os.Getenv("STREAM_CONFIG_PATH")
	if path == "" {
		return errors.New("STREAM_CONFIG_PATH not set")
	}
	streams, err := control.LoadStreams(path)
	if err != nil {
		return err
	}
	return control.Run(context.Background(), streams)
}

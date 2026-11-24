package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tedtasman/hyperstream/internal/sink"
	"github.com/tedtasman/hyperstream/internal/transform"
)

func main() {
	logCfg, _ := json.Marshal(&sink.LogConfig{})
	logger, err := sink.NewLogSink(logCfg)
	if err != nil {
		fmt.Printf("Failed to create new Log Sink: %s\n", err)
		os.Exit(1)
	}

	iCfg, _ := json.Marshal(&transform.IdentityConfig{})
	identity, err := transform.NewIdentityTransform(iCfg)

	payload := []byte("Hello, world")

	payload2, err := identity.Apply(payload)
	if err != nil {
		fmt.Printf("Failed to apply identity transform: %s\n", err)
		os.Exit(1)
	}

	err = logger.Write(payload2)
	if err != nil {
		fmt.Printf("Failed to write payload: %s\n", err)
		os.Exit(1)
	}
}

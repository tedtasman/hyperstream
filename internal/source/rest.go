package source

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tedtasman/hyperstream/internal/contract"
)

type RestConfig struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}

// implements contract.Source
type RestSource struct{ cfg RestConfig }

func NewRestSource(raw json.RawMessage) (contract.Source, error) {
	var cfg RestConfig
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return nil, err
	}
	return &RestSource{cfg: cfg}, nil
}

func (r *RestSource) Read(ctx context.Context) ([]byte, error) {
	var resp *http.Response
	var err error

	switch r.cfg.Method {
	case "GET":
		resp, err = http.Get(r.cfg.URL)
	default:
		err = fmt.Errorf("rest source config: invalid method \"%s\"", r.cfg.Method)
	}

	if err != nil {
		return nil, fmt.Errorf("rest source config: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("rest source config: %w", err)
	}

	return body, nil
}

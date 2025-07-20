//go:build js && wasm

package config

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func LoadYAML(fileName string) (*Config, error) {
	cfg := Default()
	path := filepath.Join("configs", fileName+".yaml")

	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	data, err := io.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parse config %q: %w", path, err)
	}

	return cfg, nil
}

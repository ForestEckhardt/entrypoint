package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/paketo-buildpacks/packit"
)

type EntrypointTOMLParser struct{}

func NewEntrypointTOMLParser() EntrypointTOMLParser {
	return EntrypointTOMLParser{}
}

func (e EntrypointTOMLParser) Parse(path string) ([]packit.Process, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read entrypoint.toml: %w", err)
	}

	var processes struct {
		Processes []packit.Process `toml:"processes"`
	}

	_, err = toml.DecodeReader(file, &processes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode entrypoint.toml: %w", err)
	}

	return processes.Processes, nil
}

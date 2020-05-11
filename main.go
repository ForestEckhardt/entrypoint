package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cloudfoundry/packit"
)

func main() {
	switch filepath.Base(os.Args[0]) {
	case "detect":
		packit.Detect(Detect())

	case "build":
		entrypointTOMLParser := NewEntrypointTOMLParser()

		packit.Build(Build(entrypointTOMLParser))
	default:
		panic(fmt.Sprintf("unknown command: %s", filepath.Base(os.Args[0])))
	}
}

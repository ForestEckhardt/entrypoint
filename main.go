package main

import (
	"github.com/paketo-buildpacks/packit"
)

func main() {
	entrypointTOMLParser := NewEntrypointTOMLParser()

	packit.Run(Detect(), Build(entrypointTOMLParser))
}

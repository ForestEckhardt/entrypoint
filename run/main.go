package main

import (
	"github.com/ForestEckhardt/entrypoint"
	"github.com/paketo-buildpacks/packit"
)

func main() {
	entrypointTOMLParser := entrypoint.NewEntrypointTOMLParser()

	packit.Run(
		entrypoint.Detect(),
		entrypoint.Build(entrypointTOMLParser),
	)
}

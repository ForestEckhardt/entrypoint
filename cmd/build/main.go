package main

import (
	"github.com/ForestEckhardt/entrypoint/entrypoint"

	"github.com/cloudfoundry/packit"
)

func main() {
	entrypointTOMLParser := entrypoint.NewEntrypointTOMLParser()

	packit.Build(entrypoint.Build(entrypointTOMLParser))
}

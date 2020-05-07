package main

import (
	"github.com/ForestEckhardt/entrypoint/entrypoint"

	"github.com/cloudfoundry/packit"
)

func main() {
	packit.Detect(entrypoint.Detect())
}

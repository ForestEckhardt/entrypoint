package entrypoint

import (
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

//go:generate faux --interface EntrypointParser --output fakes/entrypoint_parser.go
type EntrypointParser interface {
	Parse(path string) ([]packit.Process, error)
}

func Build(entrypointParser EntrypointParser) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		processes, err := entrypointParser.Parse(filepath.Join(context.WorkingDir, "entrypoint.toml"))
		if err != nil {
			return packit.BuildResult{}, err
		}

		return packit.BuildResult{
			Processes: processes,
		}, nil
	}
}

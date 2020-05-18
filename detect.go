package main

import (
	"os"
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		_, err := os.Stat(filepath.Join(context.WorkingDir, "entrypoint.toml"))
		if err != nil {
			return packit.DetectResult{}, packit.Fail
		}

		return packit.DetectResult{}, nil
	}
}

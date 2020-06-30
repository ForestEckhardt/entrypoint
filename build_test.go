package entrypoint_test

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/ForestEckhardt/entrypoint/fakes"
	"github.com/paketo-buildpacks/packit"
	"github.com/sclevine/spec"

	"github.com/ForestEckhardt/entrypoint"

	. "github.com/onsi/gomega"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		entrypointParser *fakes.EntrypointParser

		build packit.BuildFunc
	)

	it.Before(func() {
		entrypointParser = &fakes.EntrypointParser{}

		build = entrypoint.Build(entrypointParser)
	})

	context("there is an entrypoint.toml in the app dir", func() {
		it.Before(func() {
			entrypointParser.ParseCall.Returns.ProcessSlice = []packit.Process{
				{
					Type:    "web",
					Command: "some-start-command",
				},
			}
		})

		it("returns a result that sets the entrypoint", func() {
			result, err := build(packit.BuildContext{
				WorkingDir: "working-dir/",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Processes).To(Equal([]packit.Process{
				{
					Type:    "web",
					Command: "some-start-command",
				},
			}))
			Expect(entrypointParser.ParseCall.Receives.Path).To(Equal(filepath.Join("working-dir/entrypoint.toml")))
		})
	})

	context("failure cases", func() {
		it.Before(func() {
			entrypointParser.ParseCall.Returns.Error = errors.New("failed to parse")
		})

		it("returns a result that sets the entrypoint", func() {
			_, err := build(packit.BuildContext{})
			Expect(err).To(MatchError("failed to parse"))
		})

	})
}

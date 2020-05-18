package main_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/packit"

	main "github.com/ForestEckhardt/entrypoint"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testEntrypointTOMLParser(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect     = NewWithT(t).Expect
		workingDir string

		entrypointTOMLParser main.EntrypointTOMLParser
	)

	it.Before(func() {
		var err error
		workingDir, err = ioutil.TempDir("", "working-dir")
		Expect(err).NotTo(HaveOccurred())

		entrypointTOMLParser = main.NewEntrypointTOMLParser()
	})

	it.After(func() {
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	context("Parse", func() {
		it.Before(func() {
			Expect(ioutil.WriteFile(filepath.Join(workingDir, "entrypoint.toml"), []byte(`
[[processes]]
type = "web"
command = "some-command"
args = ["some-arg"]
direct = false
`), os.ModePerm)).To(Succeed())
		})

		it("returns a list of strings", func() {
			entrypoint, err := entrypointTOMLParser.Parse(filepath.Join(workingDir, "entrypoint.toml"))
			Expect(err).NotTo(HaveOccurred())
			Expect(entrypoint).To(Equal([]packit.Process{
				{
					Type:    "web",
					Command: "some-command",
					Args:    []string{"some-arg"},
					Direct:  false,
				},
			}))
		})

		context("failure cases", func() {
			context("when the entrypoint.toml cannot be opened", func() {
				it.Before(func() {
					Expect(os.Chmod(filepath.Join(workingDir, "entrypoint.toml"), 0000)).To(Succeed())
				})

				it("returns an error", func() {
					_, err := entrypointTOMLParser.Parse(filepath.Join(workingDir, "entrypoint.toml"))
					Expect(err).To(MatchError(ContainSubstring("failed to read entrypoint.toml:")))
					Expect(err).To(MatchError(ContainSubstring("permission denied")))
				})
			})

			context("the entrypoint.toml is malformed", func() {
				it.Before(func() {
					Expect(ioutil.WriteFile(filepath.Join(workingDir, "entrypoint.toml"), []byte("%%%"), os.ModePerm)).To(Succeed())
				})

				it("returns an error", func() {
					_, err := entrypointTOMLParser.Parse(filepath.Join(workingDir, "entrypoint.toml"))
					Expect(err).To(MatchError(ContainSubstring("failed to decode entrypoint.toml:")))
					Expect(err).To(MatchError(ContainSubstring("bare keys cannot contain '%'")))
				})
			})
		})
	})
}

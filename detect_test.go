package main_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudfoundry/packit"
	"github.com/sclevine/spec"

	main "github.com/ForestEckhardt/entrypoint"

	. "github.com/onsi/gomega"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		workingDir string

		detect packit.DetectFunc
	)

	it.Before(func() {
		var err error
		workingDir, err = ioutil.TempDir("", "workingDir")
		Expect(err).NotTo(HaveOccurred())

		detect = main.Detect()
	})

	it.After(func() {
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	context("when there is a entrypoint.toml", func() {
		it.Before(func() {
			Expect(ioutil.WriteFile(filepath.Join(workingDir, "entrypoint.toml"), []byte(``), 0644)).To(Succeed())
		})

		it("returns a plan that provides and requires entrypoint", func() {
			result, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Plan).To(Equal(packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: "entrypoint"},
				},
				Requires: []packit.BuildPlanRequirement{
					{
						Name: "entrypoint",
					},
				},
			}))
		})
	})

	context("failure cases", func() {
		context("entrypoint.toml does not exist", func() {
			it("fails detection", func() {
				_, err := detect(packit.DetectContext{
					WorkingDir: workingDir,
				})
				Expect(err).To(MatchError(packit.Fail))
			})
		})
	})
}

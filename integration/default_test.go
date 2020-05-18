package integration_test

import (
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testDefault(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		pack   occam.Pack
		docker occam.Docker

		image     occam.Image
		container occam.Container
		name      string
	)

	it.Before(func() {
		pack = occam.NewPack()
		docker = occam.NewDocker()

		var err error
		name, err = occam.RandomName()
		Expect(err).NotTo(HaveOccurred())
	})

	it.After(func() {
		Expect(docker.Container.Remove.Execute(container.ID)).To(Succeed())
		Expect(docker.Image.Remove.Execute(image.ID)).To(Succeed())
		Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())
	})

	it("should build a working OCI image for a simple start command ", func() {
		var err error
		image, _, err = pack.WithNoColor().Build.
			WithNoPull().
			WithBuildpacks(buildpack).
			Execute(name, filepath.Join("testdata", "simple-start"))
		Expect(err).ToNot(HaveOccurred())

		container, err = docker.Container.Run.Execute(image.ID)
		Expect(err).NotTo(HaveOccurred())

		logs, err := docker.Container.Logs.Execute(container.ID)
		Expect(err).NotTo(HaveOccurred())

		Expect(logs).To(MatchRegexp(`This is from the entrypoint start command`), ContainerLogs(container.ID))
	})
}

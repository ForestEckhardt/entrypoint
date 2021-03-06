package entrypoint_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitEntrypoint(t *testing.T) {
	suite := spec.New("entrypoint", spec.Report(report.Terminal{}))
	suite("Build", testBuild)
	suite("Detect", testDetect)
	suite("EntrypointTOMLParser", testEntrypointTOMLParser)
	suite.Run(t)
}

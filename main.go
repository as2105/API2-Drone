package main

import (
	"github.com/SynapticHealthAlliance/fhir-api/cmd"
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/logging"
	"github.com/SynapticHealthAlliance/fhir-api/internal/pkg/metadata"
)

var (
	commit, branch, version, buildTime string
	log                                = logging.NewLogger()
)

func main() {
	metadata.Data.Commit = commit
	metadata.Data.Branch = branch
	metadata.Data.Version = version
	if err := metadata.Data.SetBuildTime(buildTime); err != nil {
		log.WithError(err).Warn("could not parse build time")
	}

	cmd.Execute()
}

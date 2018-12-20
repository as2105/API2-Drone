package main

import (
	"github.com/SynapticHealthAlliance/fhir-api/cmd"
	"github.com/SynapticHealthAlliance/fhir-api/internal/metadata"
	"github.com/SynapticHealthAlliance/fhir-api/logging"
)

var (
	commit, branch, version, buildTime string
	log                                = logging.NewLogger()
)

func main() {
	metadata.Metadata.Commit = commit
	metadata.Metadata.Branch = branch
	metadata.Metadata.Version = version
	if err := metadata.Metadata.SetBuildTime(buildTime); err != nil {
		log.WithError(err).Warn("could not parse build time")
	}

	cmd.Execute()
}

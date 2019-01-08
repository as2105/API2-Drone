package metadata

import (
	"time"

	"github.com/pkg/errors"
)

type data struct {
	AppName   string     `json:"app_name,omitempty"`
	Commit    string     `json:"commit,omitempty"`
	Branch    string     `json:"branch,omitempty"`
	BuildTime *time.Time `json:"build_time,omitempty"`
	Version   string     `json:"version,omitempty"`
}

func (d *data) SetBuildTime(s string) error {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return errors.Wrap(err, "could not parse build time")
	}
	d.BuildTime = &t
	return nil
}

// Data is a global containing the application's metadata
var Data = data{AppName: "PDX FHIR API"}

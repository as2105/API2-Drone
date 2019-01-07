package metadata

import (
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type MetadataType struct {
	AppName   string `json:"app_name,omitempty"`
	Commit    string `json:"commit,omitempty"`
	Branch    string `json:"branch,omitempty"`
	BuildTime string `json:"build_time,omitempty"`
	Version   string `json:"version,omitempty"`
}

func (m *MetadataType) SetBuildTime(s string) error {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return errors.Wrap(err, "could not parse build time")
	}
	m.BuildTime = t.Format(http.TimeFormat)
	return nil
}

var Metadata = MetadataType{
	AppName: "PDX FHIR API",
}

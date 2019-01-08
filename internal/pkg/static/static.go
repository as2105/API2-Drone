package static

import "github.com/gobuffalo/packr/v2"

// NewStaticFilesBox creates a new Packr box containing static files
func NewStaticFilesBox() *packr.Box {
	return packr.New("Static Files", "./files")
}

package static

import "github.com/gobuffalo/packr/v2"

func NewStaticFilesBox() *packr.Box {
	return packr.New("Static Files", "./files")
}

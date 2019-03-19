package transcoder

import (
	"fmt"

	cfmt "github.com/AlbinoDrought/creamy-transcode/format"
)

type specialContainerHandler func(format *cfmt.Format) [][]string

var specialContainerHandlers = map[string]specialContainerHandler{
	"jpg": dumpSingleJPG,
	"png": dumpSinglePNG,
}

func dumpSingleThumbnail(base []string, format *cfmt.Format) []string {
	base = append(base, "-y", "-hide_banner", "-frames:v", "1")

	vf := "thumbnail"
	if format.VideoSpecs.ResolutionHeight != 0 || format.VideoSpecs.ResolutionWidth != 0 {
		height := format.VideoSpecs.ResolutionHeight
		width := format.VideoSpecs.ResolutionWidth

		// -1 signifies automatic in ffmpeg
		if height == 0 {
			height = -1
		} else if width == 0 {
			width = -1
		}

		vf += fmt.Sprintf("scale=%d:%d", width, height)
	}

	base = append(base, "-vf", vf)

	return base
}

func dumpSingleJPG(format *cfmt.Format) [][]string {
	return [][]string{
		dumpSingleThumbnail(
			[]string{"-f", "singlejpeg"},
			format,
		),
	}
}

func dumpSinglePNG(format *cfmt.Format) [][]string {
	return [][]string{
		dumpSingleThumbnail(
			[]string{"-c:v", "png"},
			format,
		),
	}
}

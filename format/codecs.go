package format

import (
	"regexp"
	"strings"
)

var supportedCodecs = []string{
	"mpeg4",
	"xvid",
	"flv",
	"h263",
	"mjpeg",
	"mpeg1video",
	"mpeg2video",
	"qtrle",
	"svq3",
	"wmv1",
	"wmv2",
	"huffyuv",
	"rv20",
	"h264",
	"hevc",
	"vp8",
	"vp9",
	"theora",
	"dnxhd",
}

var videoCodecRegex = regexp.MustCompile("^(" + strings.Join(supportedCodecs, "|") + ")$")

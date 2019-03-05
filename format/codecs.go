package format

import (
	"regexp"
	"strings"
)

var supportedVideoCodecs = []string{
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

var videoCodecRegex = regexp.MustCompile("^(" + strings.Join(supportedVideoCodecs, "|") + ")$")

var supportedAudioCodecs = []string{
	"mp3",
	"mp2",
	"aac",
	"amr_nb",
	"ac3",
	"vorbis",
	"flac",
	"pcm_u8",
	"pcm_s16le",
	"pcm_alaw",
	"wmav2",
}

var audioCodecRegex = regexp.MustCompile("^(" + strings.Join(supportedAudioCodecs, "|") + ")$")

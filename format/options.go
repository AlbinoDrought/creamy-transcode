package format

import (
	"regexp"
	"strings"
)

const formatOptionTwoPass = "2pass"

// const bitrateSuffix = "k"
const videoBitrateMaximum = 200000
const audioBitrateMaximum = 512000

var bitrateRegex = regexp.MustCompile("^(\\d{1,6})k$")

const fpsAutomatic = "0"
const fpsSuffix = "fps"

var resolutionRegex = regexp.MustCompile("^(\\d{1,4})x(\\d{1,4})$")

var definitionRegex = regexp.MustCompile("^(\\d{3,4})p$")

var fpsValid = []string{
	fpsAutomatic,
	"15",
	"23.98",
	"25",
	"29.97",
	"30",
}
var fpsRegex = regexp.MustCompile("^(" + strings.Join(fpsValid, "|") + ")fps$")

var sampleRateValid = []string{
	"8000",
	"11025",
	"16000",
	"22000",
	"22050",
	"24000",
	"32000",
	"44000",
	"44100",
	"48000",
}
var sampleRateRegex = regexp.MustCompile("^(" + strings.Join(sampleRateValid, "|") + ")hz$")

const audioChannelMono = "mono"
const audioChannelStereo = "stereo"

var audioChannelRegex = regexp.MustCompile("^(" + audioChannelMono + "|" + audioChannelStereo + ")$")
